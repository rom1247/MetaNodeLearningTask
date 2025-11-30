// SPDX-License-Identifier: MIT
pragma solidity ~0.8;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract BeggingContract is ERC721, Ownable {
    string private _name;
    string private _symbol;

    uint256 public startTime;
    uint256 private rankingListSize = 3;

    mapping(address => uint256) public donateAmounts; //捐赠总金额
    mapping(address => uint256) public donateTimes; //捐赠次数
    address[] public donors; //捐赠者列表

    constructor(
        string memory name_,
        string memory symbol_
    ) ERC721(name_, symbol_) Ownable(msg.sender) {}

    event Donation(address indexed account, uint256 amount);

    function startWithTimestamp() external {
        startTime = block.timestamp;
    }

    function start(uint256 startTime_) external {
        startTime = startTime_;
    }

    function donate() external payable returns (bool) {
        require(
            block.timestamp >= startTime,
            "The donation hasn't started yet"
        );
        require(msg.value > 0, " need greater than 0 ETH");

        address from = msg.sender;
        uint256 amount = donateAmounts[from];

        //当前捐赠为0的时候，说明是第一次捐赠
        if (amount == 0) {
            donors.push(from);
        }
        donateAmounts[from] += msg.value;
        donateTimes[from] += 1;
        emit Donation(from, msg.value);
        return true;
    }

    function withdraw() external payable onlyOwner returns (bool) {
        uint256 balance = address(this).balance;
        require(balance > 0, " 0 eth");
        (bool success, ) = owner().call{value: balance}("");

        require(success, "Transfer failed");
        return true;
    }

    function getDonation(address account) external view returns (uint256) {
        return donateAmounts[account];
    }

    function getlent() external view returns (uint256) {
        uint256 len = donors.length;
        uint256 top = rankingListSize;
        if (top > len) {
            top = len;
        }
        return top;
    }

    function getRankingList()
        external
        view
        returns (address[] memory topAddresses, uint256[] memory topAmounts)
    {
        uint256 len = donors.length;
        uint256 top = rankingListSize;
        if (top > len) {
            top = len;
        }
        topAddresses = new address[](len);
        topAmounts = new uint256[](len);

        for (uint256 i = 0; i < len; i++) {
            address d = donors[i];
            uint256 amount = donateAmounts[d];

            // 将该项尝试放入 top
            for (uint256 j = 0; j < top; j++) {
                if (amount > topAmounts[j]) {
                    //  插入排序方式，把 j 位置及后面的往后移
                    for (uint256 k = (top - 1); k > j; k--) {
                        topAmounts[k] = topAmounts[k - 1];
                        topAddresses[k] = topAddresses[k - 1];
                    }
                    topAmounts[j] = amount;
                    topAddresses[j] = d;
                    break;
                }
            }
        }
    }
}
