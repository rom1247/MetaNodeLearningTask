// SPDX-License-Identifier: MIT
pragma solidity ~0.8;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

// 一个发布的合约地址 https://sepolia.etherscan.io/token/0x468D93eB70BdaFABCDd96ECc64539754B18c0f50
contract ERC721NTF is ERC721URIStorage, Ownable {
    //设置一个计算器，使用using使用把Counters的function加载到结构体Counter中，之后就可以使用Counter.function,
    using Counters for Counters.Counter;
    using Strings for uint256;
    Counters.Counter private _tokenIds;

    string private _baseTokenURI;

    constructor(
        string memory name_,
        string memory symbol_,
        string memory baseTokenURI_
    ) ERC721(name_, symbol_) Ownable(msg.sender) {
        _baseTokenURI = baseTokenURI_;
    }

    // 设置/更新 baseURI（仅 owner）
    function setBaseURI(string memory baseURI_) external onlyOwner {
        _baseTokenURI = baseURI_;
    }

    // 返回 baseURI（OpenZeppelin 会在 tokenURI 中调用它）
    //  tokenURL有两种组装方式一种是baseURI + tokenId.json，一种是全路径tokenURI
    // 重写这个方法 就可以实现 baseURI + tokenId.json
    function _baseURI() internal view override returns (string memory) {
        return _baseTokenURI;
    }

    // 自动拼接 tokenId.json
    function tokenURI(
        uint256 tokenId
    ) public view override returns (string memory) {
        // 默认自动拼接： baseURI + tokenId + ".json"
        return
            string(abi.encodePacked(_baseURI(), tokenId.toString(), ".json"));
    }

    //   ipfs://bafybeidozh3wsjnio7gw5lw3sftkf6tv6m2ylnuquk6ac6q3eofgpzllcy/
    //    https://gateway.pinata.cloud/ipfs/bafybeidozh3wsjnio7gw5lw3sftkf6tv6m2ylnuquk6ac6q3eofgpzllcy/1.json

    // 铸造一个 token 自增 ID
    function mint(address to) external onlyOwner returns (uint256) {
        _tokenIds.increment();
        uint256 newId = _tokenIds.current();
        _safeMint(to, newId);
        return newId;
    }
}
