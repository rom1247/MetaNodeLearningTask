## transferFrom接口如何使用

transferFrom是授权转账接口，要使用要先授权（调用approve接口）或者自己转账
ERC20场景
```solitidy
# B用户授权给合约A 100个代币，然后合约A再转账给C 100个代币 
token.approve(address(A), 100);
IERC20(token).transferFrom(B, C, 100);

# 自己转账
IERC20(token).transferFrom(msg.sender, address(this), amount);
```
ERC721和ERC20的transferFrom接口不同的区别是 ERC721没有数量参数，而是tokenId


## ERC721中的授权接口跟ERC20有何不同
ERC721授权的叫NFT(non fungible token)，即非同质化代币，意思某一资产的唯一标识符，tokenId本身没有价值，具体价值由tokenId关联的metadata决定
ERC20授权是账号的余额，这种是同质化代币，无法区分资产

## safeTransfer等带safe前缀的接口提供了什么安全措施
ERC721要求目标合约实现IERC721Receiver接口的onERC721Received函数，
而ERC721会利用_checkOnERC721Received检查目标合约实现了onERC721Received()函数，
校验是否实现并返回IERC721Receiver.onERC721Received.selector，如果不是则会revert 