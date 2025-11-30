// SPDX-License-Identifier: MIT
pragma solidity ~0.8;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

contract ERC721NTF is ERC721URIStorage {
    //设置一个计算器，使用using使用把Counters的function加载到结构体Counter中，之后就可以使用Counter.function,
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;

    constructor(
        string memory name_,
        string memory symbol_
    ) ERC721(name_, symbol_) {}

    function mintNFT(string calldata tokenURI_) external returns (uint256) {
        _tokenIds.increment();
        uint256 tokenId = _tokenIds.current();

        // Mint to caller
        _safeMint(msg.sender, tokenId);

        // Store tokenURI (ERC721URIStorage)
        _setTokenURI(tokenId, tokenURI_);

        return tokenId;
    }
}
