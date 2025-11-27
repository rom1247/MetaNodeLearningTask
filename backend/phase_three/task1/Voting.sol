// SPDX-License-Identifier: MIT
pragma solidity ~0.8;

contract Voting {
    mapping(address => uint) votes;
    address[] voters;

    function vote(address addr) public {
        uint count = votes[addr];
        if (count == 0) {
            voters.push(msg.sender);
        }
        count++;
        votes[addr] = count;
    }

    function getVote(address addr) public view returns (uint) {
        return votes[addr];
    }

    function resetVotes() public {
        for (uint i = 0; i < voters.length; i++) {
            votes[voters[i]] = 0;
        }
        delete voters;
    }
}
