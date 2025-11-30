// SPDX-License-Identifier: MIT
pragma solidity ~0.8;

contract ERC20 {
    uint private constant TOTAL_SUPPLY = 1000;

    mapping(address => uint256) public _balance;
    mapping(address => mapping(address => uint256)) public _allowances;

    string public name;
    string public symbol;
    uint256 public totalSuppley;
    address public owner;

    constructor(string memory name_, string memory symbol_) {
        name = name_;
        symbol = symbol_;
        owner = msg.sender;
    }

    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(
        address indexed owner,
        address indexed sepnder,
        uint256 value
    );

    modifier onlyOwner() {
        require(owner == msg.sender, "not owner");
        _;
    }

    function balanceOf(address account) external view returns (uint256) {
        return _balance[account];
    }
    function transfer(address to, uint256 value) external returns (bool) {
        address from = msg.sender;
        uint256 fromBalance = _balance[from];
        require(fromBalance >= value, "Insufficient balance");
        _balance[from] = fromBalance - value;
        _balance[to] += value;
        emit Transfer(from, to, value);
        return true;
    }
    function approve(address spender, uint256 value) external returns (bool) {
        address from = msg.sender;
        uint256 fromBalance = _balance[from];
        require(fromBalance >= value, "Insufficient limit");
        _allowances[from][spender] += value;
        emit Approval(from, spender, value);
        return true;
    }
    function transferFrom(
        address from,
        address to,
        uint256 value
    ) external payable returns (bool) {
        //这里要检查的是被授权地址是否有代扣额度
        uint256 _allowance = _allowances[from][msg.sender];
        require(_allowance >= value, "Insufficient allowances");
        _allowances[from][to] -= _allowance - value;
        uint256 fromBalance = _balance[from];
        require(fromBalance >= value, "Insufficient balance");
        _balance[from] = fromBalance - value;
        _balance[to] += value;
        emit Transfer(from, to, value);
        return true;
    }

    function mint(address to, uint256 value) external onlyOwner {
        require(value > 0, "mint value must be greater than zero");
        totalSuppley += value;
        require(totalSuppley <= TOTAL_SUPPLY, "exceeding the total suppley");
        _balance[to] += value;
        emit Transfer(address(0), to, value);
    }
}
