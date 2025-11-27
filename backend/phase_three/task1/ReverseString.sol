// SPDX-License-Identifier: MIT
pragma solidity ~0.8;


contract ReverseString {

      function reverse(string memory str) public pure returns (string memory){
            bytes memory b = bytes(str);
            bytes memory r= new bytes(b.length);

            for (uint i = 0 ; i < b.length; i++) {
                r[i] = b[b.length -1 -i];
            }
            return string(r);

      }

}