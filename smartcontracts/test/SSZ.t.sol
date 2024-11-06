// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {Test, console} from "forge-std/Test.sol";
import {SSZHelper} from "../src/SSZHelper.sol";

contract SSZHelperTest is Test {
    using SSZHelper for bytes32;

    struct TestCase {
        bytes32 littleEndian;
        uint256 expected;
    }

    function setUp() public {}

    function test_littleEndianToUint256() public pure {
        uint256 numberOfTestCases = 10; // Replace with the actual number of test cases
        TestCase[] memory inputs = new TestCase[](numberOfTestCases);

        // Test cases generated via formula in https://github.com/ethereum/consensus-specs/blob/dev/ssz/simple-serialize.md#uintn

        inputs[0] = TestCase(0x0000000000000000000000000000000000000000000000000000000000000000, 0);

        inputs[1] = TestCase(0x0100000000000000000000000000000000000000000000000000000000000000, 1);

        inputs[2] = TestCase(
            0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff,
            115792089237316195423570985008687907853269984665640564039457584007913129639935
        );
        inputs[3] = TestCase(
            0x811c9ff8ab972f2bebf077cc25cdabf13a453220c18c5ec46311b15658f445ef,
            108226369666583997551138325543566419976152426281316995842230964140211288611969
        );
        inputs[4] = TestCase(
            0xe72e6ddd7893a1615471ab06829600553d8c991aa80327046346578fa50c8d67,
            46837436124653112299468049186755363737589723552891482908895215055147558055655
        );

        inputs[5] = TestCase(
            0x0000000000000000000000000000000000000000000000000000000000000080,
            57896044618658097711785492504343953926634992332820282019728792003956564819968
        );

        inputs[6] = TestCase(
            0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7f,
            57896044618658097711785492504343953926634992332820282019728792003956564819967
        );

        inputs[7] = TestCase(
            0x0016247e239624a7911b195077d6c7778041093f2d09994494b722e51ca4819c,
            70789860315756145078625863226032425071399523783678674187429300395657522320896
        );

        inputs[8] = TestCase(
            0x5ea8958720cb3530b302a8fec6d56c2b5e3187a42b0a796277b9f463cb6e1abc,
            85081518232832230086462199778038298743507557190517623229822604643623599712350
        );
        inputs[9] =
            TestCase(0xf74fca53e576089a2b6b00000000000000000000000000000000000000000000, 506097522914230528528375);

        for (uint256 i = 0; i < inputs.length; i++) {
            TestCase memory testCase = inputs[i];
            // Perform your test using testCase.le and testCase.expected
            uint256 result = testCase.littleEndian.fromLittleEndian();
            assertEq(result, testCase.expected);
        }
    }
}
