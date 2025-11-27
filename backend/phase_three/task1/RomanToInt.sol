// SPDX-License-Identifier: MIT
pragma solidity ~0.8;

contract RomanToInt {
    mapping(bytes1 => uint) romanInt;

    constructor() {
        romanInt["I"] = 1;
        romanInt["V"] = 5;
        romanInt["X"] = 10;
        romanInt["L"] = 50;
        romanInt["C"] = 100;
        romanInt["D"] = 500;
        romanInt["M"] = 1000;
    }

    function romanToInt(string memory roman) public view returns (uint256) {
        bytes memory r = bytes(roman);

        uint n = 0;
        uint p = 0;

        for (uint i = 0; i < r.length; i++) {
            bytes1 b = r[i];
            uint num = romanInt[b];
            require(num != 0, "Illegal Roman numerals");

            // 如果当前字符比左侧（上一个字符）大或者相等，那么当前字符减去上一个字符 * 2 ，反之相加
            if (num > p && p !=0) {
                n = n + num - p - p;
            } else {
                n += num;
            }
            p = num;
        }

        return n;
    }
}
