// SPDX-License-Identifier: MIT
pragma solidity ~0.8;

contract MergeSortedArray {
    function mergeSortedArray(
        uint[] memory a,
        uint[] memory b
    ) public pure returns (uint[] memory) {
        uint i = 0;
        uint j = 0;
        uint k = 0;

        uint[] memory result = new uint[](a.length + b.length);

        while (i < a.length && j < b.length) {
            if (a[i] <= b[j]) {
                result[k++] = a[i++];
            } else {
                result[k++] = b[j++];
            }
        }

        while (i < a.length) {
            result[k++] = a[i++];
        }

        while (j < b.length) {
            result[k++] = b[j++];
        }

        return result;
    }
}
