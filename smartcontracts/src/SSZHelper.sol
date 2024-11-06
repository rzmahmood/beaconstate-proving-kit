// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/**
 * @title SSZHelper
 * @notice Provides functions to handle Simple Serialize (SSZ) operations.
 * @dev Includes a function to convert a bytes32 little-endian value to uint256 big-endian format.
 */
library SSZHelper {
    /**
     * @notice Converts a 32-byte little-endian value to a uint256 in big-endian format.
     * @dev Inspired by:
     *      https://github.com/succinctlabs/telepathy-contracts/blob/0f3c6812d6bda96dde6ab7bdd8f8391c47bf5d0b/src/libraries/SimpleSerialize.sol#L17
     *      This function reverses the byte order of the input `bytes32` value to match Solidity's big-endian `uint256` representation.
     *      It performs a series of bitwise operations to swap bytes at different levels.
     * @param le The 32-byte little-endian value as `bytes32`.
     * @return The `uint256` value after converting from little-endian to big-endian.
     */
    function fromLittleEndian(bytes32 le) public pure returns (uint256) {
        uint256 v = uint256(le);
        v = ((v & 0xFF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00) >> 8)
            | ((v & 0x00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF) << 8);
        v = ((v & 0xFFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000) >> 16)
            | ((v & 0x0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF) << 16);
        v = ((v & 0xFFFFFFFF00000000FFFFFFFF00000000FFFFFFFF00000000FFFFFFFF00000000) >> 32)
            | ((v & 0x00000000FFFFFFFF00000000FFFFFFFF00000000FFFFFFFF00000000FFFFFFFF) << 32);
        v = ((v & 0xFFFFFFFFFFFFFFFF0000000000000000FFFFFFFFFFFFFFFF0000000000000000) >> 64)
            | ((v & 0x0000000000000000FFFFFFFFFFFFFFFF0000000000000000FFFFFFFFFFFFFFFF) << 64);
        v = (v >> 128) | (v << 128);
        return v;
    }
}
