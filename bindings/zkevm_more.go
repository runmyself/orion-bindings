// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/orion-land/orion-bindings/solc"
)

const ZKEVMStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/ZKEVM.sol:ZKEVM\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/ZKEVM.sol:ZKEVM\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/ZKEVM.sol:ZKEVM\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1003,\"contract\":\"src/L1/ZKEVM.sol:ZKEVM\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/L1/ZKEVM.sol:ZKEVM\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1005,\"contract\":\"src/L1/ZKEVM.sol:ZKEVM\",\"label\":\"PORTAL\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_contract(OptimismPortal)1015\"},{\"astId\":1006,\"contract\":\"src/L1/ZKEVM.sol:ZKEVM\",\"label\":\"submitter\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_address\"},{\"astId\":1007,\"contract\":\"src/L1/ZKEVM.sol:ZKEVM\",\"label\":\"challenger\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_address\"},{\"astId\":1008,\"contract\":\"src/L1/ZKEVM.sol:ZKEVM\",\"label\":\"lastBatchSequenced\",\"offset\":20,\"slot\":\"103\",\"type\":\"t_uint64\"},{\"astId\":1009,\"contract\":\"src/L1/ZKEVM.sol:ZKEVM\",\"label\":\"lastL2BlockNumber\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_uint64\"},{\"astId\":1010,\"contract\":\"src/L1/ZKEVM.sol:ZKEVM\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1011,\"contract\":\"src/L1/ZKEVM.sol:ZKEVM\",\"label\":\"withdrawRoots\",\"offset\":0,\"slot\":\"106\",\"type\":\"t_mapping(t_bytes32,t_uint256)\"},{\"astId\":1012,\"contract\":\"src/L1/ZKEVM.sol:ZKEVM\",\"label\":\"storageBatchs\",\"offset\":0,\"slot\":\"107\",\"type\":\"t_mapping(t_uint64,t_struct(BatchStore)1017_storage)\"},{\"astId\":1013,\"contract\":\"src/L1/ZKEVM.sol:ZKEVM\",\"label\":\"confirmBatchIndex\",\"offset\":0,\"slot\":\"108\",\"type\":\"t_mapping(t_uint64,t_bool)\"},{\"astId\":1014,\"contract\":\"src/L1/ZKEVM.sol:ZKEVM\",\"label\":\"challenges\",\"offset\":0,\"slot\":\"109\",\"type\":\"t_mapping(t_uint64,t_struct(BatchChallenge)1016_storage)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_contract(OptimismPortal)1015\":{\"encoding\":\"inplace\",\"label\":\"contract OptimismPortal\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_mapping(t_bytes32,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_uint256\"},\"t_mapping(t_uint64,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint64 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint64\",\"value\":\"t_bool\"},\"t_mapping(t_uint64,t_struct(BatchChallenge)1016_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint64 =\u003e struct ZKEVM.BatchChallenge)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint64\",\"value\":\"t_struct(BatchChallenge)1016_storage\"},\"t_mapping(t_uint64,t_struct(BatchStore)1017_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint64 =\u003e struct ZKEVM.BatchStore)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint64\",\"value\":\"t_struct(BatchStore)1017_storage\"},\"t_struct(BatchChallenge)1016_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ZKEVM.BatchChallenge\",\"numberOfBytes\":\"128\"},\"t_struct(BatchStore)1017_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ZKEVM.BatchStore\",\"numberOfBytes\":\"160\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var ZKEVMStorageLayout = new(solc.StorageLayout)

var ZKEVMDeployedBin = "0x6080604052600436106101ac5760003560e01c80638ad9e18e116100ec578063e291c2f51161008a578063f2fde38b11610064578063f2fde38b1461062a578063f4daa2911461064a578063f71b51f314610661578063fc7e286d1461069157600080fd5b8063e291c2f51461056c578063eb1ec18f146105f4578063f02deda11461060957600080fd5b80638dc45d9a116100c65780638dc45d9a14610447578063acc1245a14610474578063c7ab203914610494578063e1e158a51461055057600080fd5b80638ad9e18e146103ce5780638d644bb7146104095780638da5cb5b1461041c57600080fd5b8063485cc95511610159578063534db0e211610133578063534db0e21461032757806359ef1120146103545780636177fd1814610374578063715018a6146103b957600080fd5b8063485cc955146102d45780634b21f578146102e75780634ff561921461030757600080fd5b80632e1a7d4d1161018a5780632e1a7d4d1461025a5780633a4b66f11461027a578063423fa8561461028257600080fd5b8063032ecb38146101b15780630ff754ea146101d357806323ec85181461022a575b600080fd5b3480156101bd57600080fd5b506101d16101cc3660046123b4565b6106be565b005b3480156101df57600080fd5b506065546102009073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561023657600080fd5b5061024a6102453660046123f1565b6107d8565b6040519015158152602001610221565b34801561026657600080fd5b506101d161027536600461240e565b610953565b6101d1610b5a565b34801561028e57600080fd5b506067546102bb9074010000000000000000000000000000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff9091168152602001610221565b6101d16102e2366004612427565b610d3d565b3480156102f357600080fd5b506101d1610302366004612460565b610f66565b34801561031357600080fd5b506101d16103223660046123f1565b611531565b34801561033357600080fd5b506067546102009073ffffffffffffffffffffffffffffffffffffffff1681565b34801561036057600080fd5b506101d161036f3660046124d5565b611580565b34801561038057600080fd5b5061024a61038f3660046123f1565b73ffffffffffffffffffffffffffffffffffffffff16600090815260696020526040902054151590565b3480156103c557600080fd5b506101d1611842565b3480156103da57600080fd5b506103fb6103e936600461240e565b606a6020526000908152604090205481565b604051908152602001610221565b6101d16104173660046123b4565b611856565b34801561042857600080fd5b5060335473ffffffffffffffffffffffffffffffffffffffff16610200565b34801561045357600080fd5b506066546102009073ffffffffffffffffffffffffffffffffffffffff1681565b34801561048057600080fd5b5061024a61048f3660046123b4565b611ca8565b3480156104a057600080fd5b506105066104af3660046123b4565b606d60205260009081526040902080546001820154600283015460039093015467ffffffffffffffff8316936801000000000000000090930473ffffffffffffffffffffffffffffffffffffffff16929060ff1685565b6040805167ffffffffffffffff909616865273ffffffffffffffffffffffffffffffffffffffff90941660208601529284019190915260608301521515608082015260a001610221565b34801561055c57600080fd5b506103fb670de0b6b3a764000081565b34801561057857600080fd5b506105c16105873660046123b4565b606b602052600090815260409020805460018201546002830154600384015460049094015467ffffffffffffffff90931693919290919085565b6040805167ffffffffffffffff90961686526020860194909452928401919091526060830152608082015260a001610221565b34801561060057600080fd5b506103fb606481565b34801561061557600080fd5b506068546102bb9067ffffffffffffffff1681565b34801561063657600080fd5b506101d16106453660046123f1565b611d15565b34801561065657600080fd5b506103fb620186a081565b34801561066d57600080fd5b5061024a61067c3660046123b4565b606c6020526000908152604090205460ff1681565b34801561069d57600080fd5b506103fb6106ac3660046123f1565b60696020526000908152604090205481565b6106c781611ca8565b156106d157600080fd5b67ffffffffffffffff81166000908152606b602052604081206004015442906106fe90620186a090612587565b1190508015610794576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43616e6e6f7420636f6e6669726d20626174636820696e73696465207468652060448201527f6368616c6c656e67652077696e646f770000000000000000000000000000000060648201526084015b60405180910390fd5b5067ffffffffffffffff166000908152606c6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b60675460009074010000000000000000000000000000000000000000900467ffffffffffffffff16810361080e57506000919050565b60665473ffffffffffffffffffffffffffffffffffffffff9081169083160361088d57606754606c906000906108689060019074010000000000000000000000000000000000000000900467ffffffffffffffff1661259f565b67ffffffffffffffff16815260208101919091526040016000205460ff161592915050565b60005b60675467ffffffffffffffff740100000000000000000000000000000000000000009091048116908216101561094a5767ffffffffffffffff81166000908152606d602052604090205473ffffffffffffffffffffffffffffffffffffffff848116680100000000000000009092041614801561092a575067ffffffffffffffff81166000908152606d602052604090206003015460ff16155b156109385750600192915050565b80610942816125c8565b915050610890565b50506000919050565b61095c336107d8565b156109c3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5573657220697320696e206368616c6c656e6765000000000000000000000000604482015260640161078b565b33600090815260696020526040902054816109dd57600080fd5b81816109f1670de0b6b3a764000083612587565b1115610b1b5760665473ffffffffffffffffffffffffffffffffffffffff1633148015610a41575060675474010000000000000000000000000000000000000000900467ffffffffffffffff1615155b15610b1b5760675474010000000000000000000000000000000000000000900467ffffffffffffffff166000908152606b60205260409020600401544290610a8d90620186a090612587565b1115610b1b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f7375626d69747465722073686f756c64207761697420626174636820746f206260448201527f6520636f6e6669726d0000000000000000000000000000000000000000000000606482015260840161078b565b81831115610b265750805b3360009081526069602052604081208054859290610b459084906125ef565b90915550610b5590503382611de8565b505050565b60665473ffffffffffffffffffffffffffffffffffffffff16610bd9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f5375626d69747465722063616e6e6f7420626520616464726573732830290000604482015260640161078b565b60665473ffffffffffffffffffffffffffffffffffffffff163314610c5a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f43616c6c6572206e6f74207375626d6974746572000000000000000000000000604482015260640161078b565b60665473ffffffffffffffffffffffffffffffffffffffff16600090815260696020526040902054670de0b6b3a764000090610c97903490612587565b1015610cff576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f446f206e6f74206861766520656e6f756768206465706f736974000000000000604482015260640161078b565b60665473ffffffffffffffffffffffffffffffffffffffff1660009081526069602052604081208054349290610d36908490612587565b9091555050565b600054610100900460ff1615808015610d5d5750600054600160ff909116105b80610d775750303b158015610d77575060005460ff166001145b610e03576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161078b565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610e6157600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610e69611e94565b60678054606680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8781169182179092557fffffffff000000000000000000000000000000000000000000000000000000009092169085161790915560009081526069602052604081208054349290610efa908490612587565b90915550508015610b5557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b60665473ffffffffffffffffffffffffffffffffffffffff16610fe5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f5375626d69747465722063616e6e6f7420626520616464726573732830290000604482015260640161078b565b60665473ffffffffffffffffffffffffffffffffffffffff163314611066576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f43616c6c6572206e6f74207375626d6974746572000000000000000000000000604482015260640161078b565b60665473ffffffffffffffffffffffffffffffffffffffff16600090815260696020526040902054670de0b6b3a764000011156110ff576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f496e73756666696369656e74207374616b696e6720616d6f756e740000000000604482015260640161078b565b606754819074010000000000000000000000000000000000000000900467ffffffffffffffff1660005b8281101561141857606361f61860008061114283611f33565b67ffffffffffffffff88166000908152606b6020526040902060030154919350915089898781811061117657611176612606565b90506020028101906111889190612635565b60600135146111f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f5072657669657720737461746520726f6f74206e6f7420657175616c00000000604482015260640161078b565b60006112388383878d8d8b81811061120d5761120d612606565b905060200281019061121f9190612635565b61122d906040810190612673565b506060949350505050565b905060006112a88b8b8981811061125157611251612606565b90506020028101906112639190612635565b611271906040810190612673565b8d8d8b81811061128357611283612606565b90506020028101906112959190612635565b6112a39060208101906123b4565b611fc6565b9050600060208351026020840120905042606a60008e8e8c8181106112cf576112cf612606565b90506020028101906112e19190612635565b608001358152602001908152602001600020819055506040518060a001604052808d8d8b81811061131457611314612606565b90506020028101906113269190612635565b6113349060208101906123b4565b67ffffffffffffffff1681526020018d8d8b81811061135557611355612606565b90506020028101906113679190612635565b6080908101358252602080830185905260408084018790524260609485015267ffffffffffffffff8e81166000908152606b8452829020865181547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016921691909117815591850151600183015584015160028201559183015160038301559190910151600490910155886113fb816125c8565b995050505050505050508080611410906126df565b915050611129565b50606780547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000067ffffffffffffffff84160217905583836114716001856125ef565b81811061148057611480612606565b90506020028101906114929190612635565b6114a09060208101906123b4565b606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff928316908117909155606754604051918252740100000000000000000000000000000000000000009004909116907faff2dd67e4faff27d6d11dbdc2eda3379b944ad1b67b973265a60efd9cd08ac9906020015b60405180910390a250505050565b611539611fd0565b606580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b67ffffffffffffffff83166000908152606d602052604090205468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff166115c557600080fd5b67ffffffffffffffff83166000908152606d602052604090206003015460ff161561164c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f4368616c6c656e676520616c72656164792066696e6973686564000000000000604482015260640161078b565b67ffffffffffffffff83166000908152606d6020526040812060020154429061167790606490612587565b119050806116c3576116be846040518060400160405280600781526020017f74696d656f757400000000000000000000000000000000000000000000000000815250612051565b6117f9565b8161172a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f496e76616c69642070726f6f6600000000000000000000000000000000000000604482015260640161078b565b67ffffffffffffffff84166000908152606b6020526040902060020154611754908490849061214e565b6117ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f50726f7665206661696c65640000000000000000000000000000000000000000604482015260640161078b565b6117f9846040518060400160405280600d81526020017f70726f7665207375636365737300000000000000000000000000000000000000815250612169565b50505067ffffffffffffffff166000908152606d6020526040902060030180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b61184a611fd0565b6118546000612266565b565b60675473ffffffffffffffffffffffffffffffffffffffff166118d5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4368616c6c656e6765722063616e6e6f74206265206164647265737328302900604482015260640161078b565b60675473ffffffffffffffffffffffffffffffffffffffff163314611956576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f43616c6c6572206e6f74206368616c6c656e6765720000000000000000000000604482015260640161078b565b67ffffffffffffffff81166000908152606b602052604081206004015490036119db576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4261746368206e6f742065786973740000000000000000000000000000000000604482015260640161078b565b67ffffffffffffffff81166000908152606d602052604090205468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff1615611a7e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f416c726561647920686173206368616c6c656e67650000000000000000000000604482015260640161078b565b67ffffffffffffffff81166000908152606b60205260408120600401544290611aab90620186a090612587565b11905080611b3b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603360248201527f43616e6e6f74206368616c6c656e6765206261746368206f757473696465207460448201527f6865206368616c6c656e67652077696e646f7700000000000000000000000000606482015260840161078b565b670de0b6b3a7640000341015611b5057600080fd5b60675473ffffffffffffffffffffffffffffffffffffffff1660009081526069602052604081208054349290611b87908490612587565b90915550506040805160a08101825267ffffffffffffffff848116808352336020808501828152348688018181524260608901908152600060808a01818152888252606d8752908b902099518a54955173ffffffffffffffffffffffffffffffffffffffff1668010000000000000000027fffffffff00000000000000000000000000000000000000000000000000000000909616991698909817939093178855516001880155905160028701559351600390950180549515157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0090961695909517909455845190815292830191909152917fff4bd2fd14b0670e131544fdc1e7c53f06371ccac4a3a8e0445a93d02d68fb05910160405180910390a25050565b67ffffffffffffffff81166000908152606d602052604081205468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff1615801590611d0f575067ffffffffffffffff82166000908152606d602052604090206003015460ff16155b92915050565b611d1d611fd0565b73ffffffffffffffffffffffffffffffffffffffff8116611dc0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161078b565b611dc981612266565b50565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b6000611e05835a84604051806020016040528060008152506122dd565b905080610b55576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a20455448207472616e736665722066616960448201527f6c65640000000000000000000000000000000000000000000000000000000000606482015260840161078b565b600054610100900460ff16611f2b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161078b565b6118546122f7565b60008061f6188311611f4c575060039261290492509050565b620493e08311611f645750600e926201107692509050565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f434952435549545f434f4e464947000000000000000000000000000000000000604482015260640161078b565b60005b9392505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314611854576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161078b565b67ffffffffffffffff82166000908152606d602090815260408083205460665473ffffffffffffffffffffffffffffffffffffffff9081168552606990935290832080546801000000000000000090920490921692670de0b6b3a764000092916120bc9084906125ef565b909155505073ffffffffffffffffffffffffffffffffffffffff811660009081526069602052604081208054670de0b6b3a764000092906120fe908490612587565b925050819055508267ffffffffffffffff167fe70d3820e244d5f71d1a6395db24f3460e8dca966edc1fd3625b6292880a877a8284604051612141929190612717565b60405180910390a2505050565b600082810361215f57506000611fc9565b5060019392505050565b67ffffffffffffffff82166000908152606d6020908152604080832080546001909101546801000000000000000090910473ffffffffffffffffffffffffffffffffffffffff1680855260699093529083208054929391928392906121cf9084906125ef565b909155505060665473ffffffffffffffffffffffffffffffffffffffff166000908152606960205260408120805483929061220b908490612587565b909155505060665460405167ffffffffffffffff8616917fe70d3820e244d5f71d1a6395db24f3460e8dca966edc1fd3625b6292880a877a916115239173ffffffffffffffffffffffffffffffffffffffff16908790612717565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600080600080845160208601878a8af19695505050505050565b600054610100900460ff1661238e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161078b565b61185433612266565b803567ffffffffffffffff811681146123af57600080fd5b919050565b6000602082840312156123c657600080fd5b611fc982612397565b73ffffffffffffffffffffffffffffffffffffffff81168114611dc957600080fd5b60006020828403121561240357600080fd5b8135611fc9816123cf565b60006020828403121561242057600080fd5b5035919050565b6000806040838503121561243a57600080fd5b8235612445816123cf565b91506020830135612455816123cf565b809150509250929050565b6000806020838503121561247357600080fd5b823567ffffffffffffffff8082111561248b57600080fd5b818501915085601f83011261249f57600080fd5b8135818111156124ae57600080fd5b8660208260051b85010111156124c357600080fd5b60209290920196919550909350505050565b6000806000604084860312156124ea57600080fd5b6124f384612397565b9250602084013567ffffffffffffffff8082111561251057600080fd5b818601915086601f83011261252457600080fd5b81358181111561253357600080fd5b87602082850101111561254557600080fd5b6020830194508093505050509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561259a5761259a612558565b500190565b600067ffffffffffffffff838116908316818110156125c0576125c0612558565b039392505050565b600067ffffffffffffffff8083168181036125e5576125e5612558565b6001019392505050565b60008282101561260157612601612558565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4183360301811261266957600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126126a857600080fd5b83018035915067ffffffffffffffff8211156126c357600080fd5b6020019150368190038213156126d857600080fd5b9250929050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361271057612710612558565b5060010190565b73ffffffffffffffffffffffffffffffffffffffff8316815260006020604081840152835180604085015260005b8181101561276157858101830151858201606001528201612745565b81811115612773576000606083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160600194935050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(ZKEVMStorageLayoutJSON), ZKEVMStorageLayout); err != nil {
		panic(err)
	}

	layouts["ZKEVM"] = ZKEVMStorageLayout
	deployedBytecodes["ZKEVM"] = ZKEVMDeployedBin
}
