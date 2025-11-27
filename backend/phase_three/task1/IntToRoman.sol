// SPDX-License-Identifier: MIT
pragma solidity ~0.8;

contract IntToRoman {
    // 千位、百位、十位、个位的罗马数字映射

    function intToRoman(uint num) public pure returns (string memory) {
        require(num > 0 || num < 3999, "Beyond the range of Roman numerals");
        bytes memory result;
        // 千位
        for (uint i = 0; i < num / 1000; i++) {
            result = bytes.concat(result, "M");
        }

        // 百位
        result = bytes.concat(result, getHundredRomanNum((num % 1000) / 100));

        // 十位
        result = bytes.concat(result, getTenRomanNum((num % 100) / 10));
        // 个位
        result = bytes.concat(result, getOneRomanNum(num % 10));

        return string(result);
    }

    function getHundredRomanNum(
        uint b
    ) private pure returns (bytes memory result) {
        if (b < 4) {
            for (uint i = 1; i <= b; i++) {
                result = bytes.concat(result, "C");
            }
        }
        if (b == 4) {
            result = bytes.concat(result, "CD");
        }

        if (b >= 5 && b < 9) {
            result = bytes.concat(result, "D");
            for (uint i = 6; i <= b; i++) {
                result = bytes.concat(result, "C");
            }
        }
        if (b == 9) {
            result = bytes.concat(result, "CM");
        }
    }

    function getTenRomanNum(uint b) private pure returns (bytes memory result) {
        if (b < 4) {
            for (uint i = 1; i <= b; i++) {
                result = bytes.concat(result, "X");
            }
        }
        if (b == 4) {
            result = bytes.concat(result, "XL");
        }

        if (b >= 5 && b < 9) {
            result = bytes.concat(result, "L");
            for (uint i = 6; i <= b; i++) {
                result = bytes.concat(result, "X");
            }
        }
        if (b == 9) {
            result = bytes.concat(result, "XC");
        }
    }

    function getOneRomanNum(uint b) private pure returns (bytes memory result) {
        if (b < 4) {
            for (uint i = 1; i <= b; i++) {
                result = bytes.concat(result, "I");
            }
        }
        if (b == 4) {
            result = bytes.concat(result, "IV");
        }

        if (b >= 5 && b < 9) {
            result = bytes.concat(result, "V");
            for (uint i = 6; i <= b; i++) {
                result = bytes.concat(result, "I");
            }
        }
        if (b == 9) {
            result = bytes.concat(result, "IX");
        }
    }
}
