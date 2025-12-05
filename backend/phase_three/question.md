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

## delegatecall 跟 call 的区别是什么
- delegatecall
  - 在当前合约的上下文执行代码
  - 使用当前合约的状态变量
  - msg.sender是原始的外部用户
- call
  - 在被调用的合约的上下文执行代码
  - 使用的是被调合约的状态变量
  - msg.sender是当前合约

## 可升级合约的执行流程是什么（user -> proxy -> implementation）
- 用户调用代理合约
- 代理合约执行fallback函数（因为代理合约没有用户的函数，所以会执行fallback函数）
- 代理合约的fallback函数会调用delegatecall的逻辑合约的函数
- 逻辑合约执行被调用的函数，并修改代理合约的状态变量（这里有个注意：逻辑合约的状态变量永远不会被修改，因为delegatecall只影响代理合约的上下文）
- 代理合约函数执行完成，并返回结果


## 代理合约上本身是有存储的，怎么避免跟逻辑合约上的存储产生冲突
### 第一种
代理合约变量声明数量和顺序和逻辑合一致，升级合约只能往后加，不能改变顺序

### 第二种
代理合约EIP-1967标准使用固定slot存储变量，而所有的业务变量声明在逻辑合约里

为什么可以这样设计变量？
因为evn是的storage存储是按slot顺序存在的，也就是说是按变量的声明顺序存储的，实际上和变量名称没有关系，
而EIP-1967标准的固定slot都是很大的slot位置（每个合约部署后会分配2的256次方个slot，所以这个固定slot位置会非常大，几乎用不到，这里还有个要点是，slot只有赋值了才会写入数据）

## 逻辑合约升级的存储冲突问题
这个和上面问题感觉是同一个问题，实际上就是storage的slot存储顺序问题，实际上和代码上的变量名称没有关系
## 可以在逻辑合约的构造函数中初始化变量吗？为什么
不会初始化变量。
因为代理合约部署的时候，没有调用逻辑合约的构造函数，
代理合约使用delegatecall执行逻辑合约的代码逻辑，改变的是代理合约的上下文变量，而逻辑合约的构造函数只会影响逻辑合约的上下文变量，
所以逻辑合约的变量不会被初始化。
