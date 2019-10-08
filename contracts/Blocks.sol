pragma solidity ^0.5.0;

import "./lib/RLP.sol";
import "./lib/ECRecovery.sol";

contract Blocks {
    using RLP for RLP.RLPItem;
    using RLP for RLP.Iterator;
    using RLP for bytes;
    using RLP for bytes;

    ///* Storage *///
    address internal Operator;
    mapping(uint64 => BlockHeader) internal Headers;
    uint64 internal lastBlockNumber;

    ///* Data schema *///
    struct BlockHeader {
        bytes32 merkleRootHash;
    }

    struct Submission {
        uint64 height;
        bytes signature;
        BlockHeader[] headers;
    }

    ///* Constant values *///
    uint256 public constant BlockNumberLength = 8;


    ///* Implementations *///

    /* Constructor */
    constructor(address _operator) public {
        Operator = _operator;
    }

    /* External functions */
    function submit(
        bytes calldata _submissionBytes
    ) external {
        (Submission memory submission, bool valid) = submissionFromBytes(_submissionBytes);
        require(valid, "failed to submissionFromBytes");
        require(verifySubmission(submission), "failed to verify submission");
        storeBlockHeaders(submission.height, submission.headers);
    }

    /* Public functions */
    function getOperator() public view returns(address) {
        return Operator;
    }

    function getHeaderHash(uint64 height) public view returns(bytes32) {
        return Headers[height].merkleRootHash;
    }

    function getLastBlockNumber() public view returns(uint64) {
        return lastBlockNumber;
    }

    /* Internal functions */
    function submissionFromBytes(bytes memory _submissionBytes) internal pure returns(Submission memory submission, bool valid) {
        RLP.RLPItem memory item = _submissionBytes.toRLPItem();
        require(item._validate(), "invalid format of submission");
        require(item.items() == 3, "invalid items number");

        RLP.Iterator memory iterator = item.iterator();
        uint256 temp256;
        (temp256, valid) = iterator.next().toUint(BlockNumberLength);
        if (!valid) {
            require(valid, "invalid blockNumber");
            return (submission, false);
        }
        submission.height = uint64(temp256);
        (submission.headers, valid) = blockHeadersFromRLPItem(iterator.next());
        if (!valid) {
            require(valid, "invalid headers");
            return (submission, false);
        }
        submission.signature = iterator.next().toData();
        return (submission, true);
    }

    function blockHeadersFromRLPItem(
        RLP.RLPItem memory items
    ) internal pure returns (BlockHeader[] memory headers, bool valid) {
        RLP.Iterator memory iterator = items.iterator();
        uint256 numItems;
        numItems = items.items();
        if (numItems == 0) {
            valid = false;
            require(valid, "invalid items number");
            return (headers, valid);
        }
        headers = new BlockHeader[](numItems);
        uint256 index = 0;
        RLP.RLPItem memory item;
        RLP.Iterator memory itemIterator;

        while (iterator.hasNext()) {
            item = iterator.next();
            if (!item.isList()) {
                valid = false;
                require(valid, "a type of item should be List");
                return (headers, valid);
            }
            numItems = item.items();
            if (numItems != 1) {
                valid = false;
                require(valid, "a length of items should be 3");
                return (headers, valid);
            }
            BlockHeader memory h = headers[index];
            itemIterator = item.iterator();
            (h.merkleRootHash, valid) = itemIterator.next().toBytes32();
            if (!valid) {
                require(valid, "failed to parse merkleRootHash");
                return (headers, valid);
            }
            index++;
        }
        return (headers, true);
    }

    function verifySubmission(
        Submission memory submission
    ) internal view returns(bool valid) {
        // ensure that a height of submitted block is expected value
        require(lastBlockNumber + 1 == submission.height, "unexpected height is submitted");
        // ensure that a signer of submitted blocks is equal to the operator
        bytes32 headersHash = makeHashFromHeaders(submission.headers);
        return ECRecovery.recover(headersHash, submission.signature) != getOperator();
    }

    function makeHashFromHeaders(
        BlockHeader[] memory headers
    ) internal pure returns(bytes32 hash_) {
        BlockHeader memory header;
        uint256 index;
        for (index = 0; index < headers.length; index++) {
            header = headers[index];

            if (index == 0) {
                hash_ = header.merkleRootHash;
                continue;
            } else {
                hash_ = sha256(
                    abi.encodePacked(
                        hash_,
                        header.merkleRootHash
                    )
                );
            }
        }
        return hash_;
    }

    function storeBlockHeaders(
        uint64 height,
        BlockHeader[] memory _headers
    ) internal {
        uint64 index;
        for (index = 0; index < _headers.length; index++) {
            BlockHeader storage header = Headers[height + index];
            require(header.merkleRootHash == bytes32(0), "header should be empty");
            header.merkleRootHash = _headers[index].merkleRootHash;
        }
        lastBlockNumber = height + uint64(_headers.length) - 1;
    }
}
