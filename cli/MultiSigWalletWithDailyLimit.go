// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cli

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MultiSigWalletWithDailyLimitABI is the input ABI used to generate the binding from.
const MultiSigWalletWithDailyLimitABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"removeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"revokeConfirmation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"confirmations\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"calcMaxWithdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dailyLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastDay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmationCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"from\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\"},{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionIds\",\"outputs\":[{\"name\":\"_transactionIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmations\",\"outputs\":[{\"name\":\"_confirmations\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transactionCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_required\",\"type\":\"uint256\"}],\"name\":\"changeRequirement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"submitTransaction\",\"outputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_dailyLimit\",\"type\":\"uint256\"}],\"name\":\"changeDailyLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OWNER_COUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"required\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"replaceOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spentToday\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"},{\"name\":\"_required\",\"type\":\"uint256\"},{\"name\":\"_dailyLimit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"dailyLimit\",\"type\":\"uint256\"}],\"name\":\"DailyLimitChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Confirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Revocation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Submission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Execution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"RequirementChange\",\"type\":\"event\"}]"

// MultiSigWalletWithDailyLimitBin is the compiled bytecode used for deploying new contracts.
var MultiSigWalletWithDailyLimitBin = "0x60806040523480156200001157600080fd5b506040516200189438038062001894833981016040908152815160208301519183015192018051909290839083906000908260328211801590620000555750818111155b80156200006157508015155b80156200006d57508115155b15156200007957600080fd5b600092505b845183101562000151576002600086858151811015156200009b57fe5b6020908102909101810151600160a060020a031682528101919091526040016000205460ff16158015620000f157508483815181101515620000d957fe5b90602001906020020151600160a060020a0316600014155b1515620000fd57600080fd5b60016002600087868151811015156200011257fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff1916911515919091179055600192909201916200007e565b8451620001669060039060208801906200017d565b505050600491909155505060065550620002119050565b828054828255906000526020600020908101928215620001d5579160200282015b82811115620001d55782518254600160a060020a031916600160a060020a039091161782556020909201916001909101906200019e565b50620001e3929150620001e7565b5090565b6200020e91905b80821115620001e3578054600160a060020a0319168155600101620001ee565b90565b61167380620002216000396000f3006080604052600436106101535763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663025e7c278114610195578063173825d9146101c957806320ea8d86146101ea5780632f54bf6e146102025780633411c81c146102375780634bc9fdc21461025b578063547415251461028257806367eeba0c146102a15780636b0c932d146102b65780637065cb48146102cb578063784547a7146102ec5780638b51d13f146103045780639ace38c21461031c578063a0e67e2b146103d7578063a8abe69a1461043c578063b5dc40c314610461578063b77bf60014610479578063ba51a6df1461048e578063c01a8c84146104a6578063c6427474146104be578063cea0862114610527578063d74f8edd1461053f578063dc8452cd14610554578063e20056e614610569578063ee22610b14610590578063f059cf2b146105a8575b60003411156101935760408051348152905133917fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c919081900360200190a25b005b3480156101a157600080fd5b506101ad6004356105bd565b60408051600160a060020a039092168252519081900360200190f35b3480156101d557600080fd5b50610193600160a060020a03600435166105e5565b3480156101f657600080fd5b5061019360043561075c565b34801561020e57600080fd5b50610223600160a060020a0360043516610816565b604080519115158252519081900360200190f35b34801561024357600080fd5b50610223600435600160a060020a036024351661082b565b34801561026757600080fd5b5061027061084b565b60408051918252519081900360200190f35b34801561028e57600080fd5b5061027060043515156024351515610885565b3480156102ad57600080fd5b506102706108f1565b3480156102c257600080fd5b506102706108f7565b3480156102d757600080fd5b50610193600160a060020a03600435166108fd565b3480156102f857600080fd5b50610223600435610a22565b34801561031057600080fd5b50610270600435610aa6565b34801561032857600080fd5b50610334600435610b15565b6040518085600160a060020a0316600160a060020a031681526020018481526020018060200183151515158152602001828103825284818151815260200191508051906020019080838360005b83811015610399578181015183820152602001610381565b50505050905090810190601f1680156103c65780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b3480156103e357600080fd5b506103ec610bd3565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015610428578181015183820152602001610410565b505050509050019250505060405180910390f35b34801561044857600080fd5b506103ec60043560243560443515156064351515610c35565b34801561046d57600080fd5b506103ec600435610d6e565b34801561048557600080fd5b50610270610ee7565b34801561049a57600080fd5b50610193600435610eed565b3480156104b257600080fd5b50610193600435610f6c565b3480156104ca57600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610270948235600160a060020a03169460248035953695946064949201919081908401838280828437509497506110379650505050505050565b34801561053357600080fd5b50610193600435611056565b34801561054b57600080fd5b5061027061109d565b34801561056057600080fd5b506102706110a2565b34801561057557600080fd5b50610193600160a060020a03600435811690602435166110a8565b34801561059c57600080fd5b50610193600435611232565b3480156105b457600080fd5b50610270611448565b60038054829081106105cb57fe5b600091825260209091200154600160a060020a0316905081565b60003330146105f357600080fd5b600160a060020a038216600090815260026020526040902054829060ff16151561061c57600080fd5b600160a060020a0383166000908152600260205260408120805460ff1916905591505b600354600019018210156106f75782600160a060020a031660038381548110151561066657fe5b600091825260209091200154600160a060020a031614156106ec5760038054600019810190811061069357fe5b60009182526020909120015460038054600160a060020a0390921691849081106106b957fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a031602179055506106f7565b60019091019061063f565b60038054600019019061070a9082611586565b5060035460045411156107235760035461072390610eed565b604051600160a060020a038416907f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b9090600090a2505050565b3360008181526002602052604090205460ff16151561077a57600080fd5b60008281526001602090815260408083203380855292529091205483919060ff1615156107a657600080fd5b600084815260208190526040902060030154849060ff16156107c757600080fd5b6000858152600160209081526040808320338085529252808320805460ff191690555187927ff6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e991a35050505050565b60026020526000908152604090205460ff1681565b600160209081526000928352604080842090915290825290205460ff1681565b600060075462015180014211156108655750600654610882565b600854600654101561087957506000610882565b50600854600654035b90565b6000805b6005548110156108ea578380156108b2575060008181526020819052604090206003015460ff16155b806108d657508280156108d6575060008181526020819052604090206003015460ff165b156108e2576001820191505b600101610889565b5092915050565b60065481565b60075481565b33301461090957600080fd5b600160a060020a038116600090815260026020526040902054819060ff161561093157600080fd5b81600160a060020a038116151561094757600080fd5b600380549050600101600454603282111580156109645750818111155b801561096f57508015155b801561097a57508115155b151561098557600080fd5b600160a060020a038516600081815260026020526040808220805460ff1916600190811790915560038054918201815583527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01805473ffffffffffffffffffffffffffffffffffffffff191684179055517ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d9190a25050505050565b600080805b600354811015610a9f5760008481526001602052604081206003805491929184908110610a5057fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff1615610a84576001820191505b600454821415610a975760019250610a9f565b600101610a27565b5050919050565b6000805b600354811015610b0f5760008381526001602052604081206003805491929184908110610ad357fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff1615610b07576001820191505b600101610aaa565b50919050565b6000602081815291815260409081902080546001808301546002808501805487516101009582161595909502600019011691909104601f8101889004880284018801909652858352600160a060020a0390931695909491929190830182828015610bc05780601f10610b9557610100808354040283529160200191610bc0565b820191906000526020600020905b815481529060010190602001808311610ba357829003601f168201915b5050506003909301549192505060ff1684565b60606003805480602002602001604051908101604052809291908181526020018280548015610c2b57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610c0d575b5050505050905090565b606080600080600554604051908082528060200260200182016040528015610c67578160200160208202803883390190505b50925060009150600090505b600554811015610cee57858015610c9c575060008181526020819052604090206003015460ff16155b80610cc05750848015610cc0575060008181526020819052604090206003015460ff165b15610ce657808383815181101515610cd457fe5b60209081029091010152600191909101905b600101610c73565b878703604051908082528060200260200182016040528015610d1a578160200160208202803883390190505b5093508790505b86811015610d63578281815181101515610d3757fe5b9060200190602002015184898303815181101515610d5157fe5b60209081029091010152600101610d21565b505050949350505050565b606080600080600380549050604051908082528060200260200182016040528015610da3578160200160208202803883390190505b50925060009150600090505b600354811015610e605760008581526001602052604081206003805491929184908110610dd857fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff1615610e58576003805482908110610e1357fe5b6000918252602090912001548351600160a060020a0390911690849084908110610e3957fe5b600160a060020a03909216602092830290910190910152600191909101905b600101610daf565b81604051908082528060200260200182016040528015610e8a578160200160208202803883390190505b509350600090505b81811015610edf578281815181101515610ea857fe5b906020019060200201518482815181101515610ec057fe5b600160a060020a03909216602092830290910190910152600101610e92565b505050919050565b60055481565b333014610ef957600080fd5b6003548160328211801590610f0e5750818111155b8015610f1957508015155b8015610f2457508115155b1515610f2f57600080fd5b60048390556040805184815290517fa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a9181900360200190a1505050565b3360008181526002602052604090205460ff161515610f8a57600080fd5b6000828152602081905260409020548290600160a060020a03161515610faf57600080fd5b60008381526001602090815260408083203380855292529091205484919060ff1615610fda57600080fd5b6000858152600160208181526040808420338086529252808420805460ff1916909317909255905187927f4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef91a361103085611232565b5050505050565b600061104484848461144e565b905061104f81610f6c565b9392505050565b33301461106257600080fd5b60068190556040805182815290517fc71bdc6afaf9b1aa90a7078191d4fc1adf3bf680fca3183697df6b0dc226bca29181900360200190a150565b603281565b60045481565b60003330146110b657600080fd5b600160a060020a038316600090815260026020526040902054839060ff1615156110df57600080fd5b600160a060020a038316600090815260026020526040902054839060ff161561110757600080fd5b600092505b6003548310156111985784600160a060020a031660038481548110151561112f57fe5b600091825260209091200154600160a060020a0316141561118d578360038481548110151561115a57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550611198565b60019092019161110c565b600160a060020a03808616600081815260026020526040808220805460ff1990811690915593881682528082208054909416600117909355915190917f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b9091a2604051600160a060020a038516907ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d90600090a25050505050565b336000818152600260205260408120549091829160ff16151561125457600080fd5b60008481526001602090815260408083203380855292529091205485919060ff16151561128057600080fd5b600086815260208190526040902060030154869060ff16156112a157600080fd5b600087815260208190526040902095506112ba87610a22565b945084806112ed57506002808701546000196101006001831615020116041580156112ed57506112ed866001015461153e565b1561143f5760038601805460ff191660011790558415156113175760018601546008805490910190555b8560000160009054906101000a9004600160a060020a0316600160a060020a031686600101548760020160405180828054600181600116156101000203166002900480156113a65780601f1061137b576101008083540402835291602001916113a6565b820191906000526020600020905b81548152906001019060200180831161138957829003601f168201915b505091505060006040518083038185875af192505050156113f15760405187907f33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed7590600090a261143f565b60405187907f526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b7923690600090a260038601805460ff1916905584151561143f576001860154600880549190910390555b50505050505050565b60085481565b600083600160a060020a038116151561146657600080fd5b60055460408051608081018252600160a060020a0388811682526020808301898152838501898152600060608601819052878152808452959095208451815473ffffffffffffffffffffffffffffffffffffffff1916941693909317835551600183015592518051949650919390926114e69260028501929101906115af565b50606091909101516003909101805460ff191691151591909117905560058054600101905560405182907fc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e5190600090a2509392505050565b60006007546201518001421115611559574260075560006008555b600654826008540111806115705750600854828101105b1561157d57506000611581565b5060015b919050565b8154818355818111156115aa576000838152602090206115aa91810190830161162d565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106115f057805160ff191683800117855561161d565b8280016001018555821561161d579182015b8281111561161d578251825591602001919060010190611602565b5061162992915061162d565b5090565b61088291905b8082111561162957600081556001016116335600a165627a7a723058201b9067f97755a5c1ae5e792a0b31c74fc78c2e97026334235d117b1b504bf6df0029"

// DeployMultiSigWalletWithDailyLimit deploys a new Ethereum contract, binding an instance of MultiSigWalletWithDailyLimit to it.
func DeployMultiSigWalletWithDailyLimit(auth *bind.TransactOpts, backend bind.ContractBackend, _owners []common.Address, _required *big.Int, _dailyLimit *big.Int) (common.Address, *types.Transaction, *MultiSigWalletWithDailyLimit, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiSigWalletWithDailyLimitABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MultiSigWalletWithDailyLimitBin), backend, _owners, _required, _dailyLimit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiSigWalletWithDailyLimit{MultiSigWalletWithDailyLimitCaller: MultiSigWalletWithDailyLimitCaller{contract: contract}, MultiSigWalletWithDailyLimitTransactor: MultiSigWalletWithDailyLimitTransactor{contract: contract}, MultiSigWalletWithDailyLimitFilterer: MultiSigWalletWithDailyLimitFilterer{contract: contract}}, nil
}

// MultiSigWalletWithDailyLimit is an auto generated Go binding around an Ethereum contract.
type MultiSigWalletWithDailyLimit struct {
	MultiSigWalletWithDailyLimitCaller     // Read-only binding to the contract
	MultiSigWalletWithDailyLimitTransactor // Write-only binding to the contract
	MultiSigWalletWithDailyLimitFilterer   // Log filterer for contract events
}

// MultiSigWalletWithDailyLimitCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiSigWalletWithDailyLimitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletWithDailyLimitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiSigWalletWithDailyLimitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletWithDailyLimitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiSigWalletWithDailyLimitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletWithDailyLimitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiSigWalletWithDailyLimitSession struct {
	Contract     *MultiSigWalletWithDailyLimit // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MultiSigWalletWithDailyLimitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiSigWalletWithDailyLimitCallerSession struct {
	Contract *MultiSigWalletWithDailyLimitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// MultiSigWalletWithDailyLimitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiSigWalletWithDailyLimitTransactorSession struct {
	Contract     *MultiSigWalletWithDailyLimitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// MultiSigWalletWithDailyLimitRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiSigWalletWithDailyLimitRaw struct {
	Contract *MultiSigWalletWithDailyLimit // Generic contract binding to access the raw methods on
}

// MultiSigWalletWithDailyLimitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiSigWalletWithDailyLimitCallerRaw struct {
	Contract *MultiSigWalletWithDailyLimitCaller // Generic read-only contract binding to access the raw methods on
}

// MultiSigWalletWithDailyLimitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiSigWalletWithDailyLimitTransactorRaw struct {
	Contract *MultiSigWalletWithDailyLimitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiSigWalletWithDailyLimit creates a new instance of MultiSigWalletWithDailyLimit, bound to a specific deployed contract.
func NewMultiSigWalletWithDailyLimit(address common.Address, backend bind.ContractBackend) (*MultiSigWalletWithDailyLimit, error) {
	contract, err := bindMultiSigWalletWithDailyLimit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimit{MultiSigWalletWithDailyLimitCaller: MultiSigWalletWithDailyLimitCaller{contract: contract}, MultiSigWalletWithDailyLimitTransactor: MultiSigWalletWithDailyLimitTransactor{contract: contract}, MultiSigWalletWithDailyLimitFilterer: MultiSigWalletWithDailyLimitFilterer{contract: contract}}, nil
}

// NewMultiSigWalletWithDailyLimitCaller creates a new read-only instance of MultiSigWalletWithDailyLimit, bound to a specific deployed contract.
func NewMultiSigWalletWithDailyLimitCaller(address common.Address, caller bind.ContractCaller) (*MultiSigWalletWithDailyLimitCaller, error) {
	contract, err := bindMultiSigWalletWithDailyLimit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitCaller{contract: contract}, nil
}

// NewMultiSigWalletWithDailyLimitTransactor creates a new write-only instance of MultiSigWalletWithDailyLimit, bound to a specific deployed contract.
func NewMultiSigWalletWithDailyLimitTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiSigWalletWithDailyLimitTransactor, error) {
	contract, err := bindMultiSigWalletWithDailyLimit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitTransactor{contract: contract}, nil
}

// NewMultiSigWalletWithDailyLimitFilterer creates a new log filterer instance of MultiSigWalletWithDailyLimit, bound to a specific deployed contract.
func NewMultiSigWalletWithDailyLimitFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiSigWalletWithDailyLimitFilterer, error) {
	contract, err := bindMultiSigWalletWithDailyLimit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitFilterer{contract: contract}, nil
}

// bindMultiSigWalletWithDailyLimit binds a generic wrapper to an already deployed contract.
func bindMultiSigWalletWithDailyLimit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiSigWalletWithDailyLimitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultiSigWalletWithDailyLimit.Contract.MultiSigWalletWithDailyLimitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.MultiSigWalletWithDailyLimitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.MultiSigWalletWithDailyLimitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultiSigWalletWithDailyLimit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.contract.Transact(opts, method, params...)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) MAXOWNERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "MAX_OWNER_COUNT")
	return *ret0, err
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.MAXOWNERCOUNT(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.MAXOWNERCOUNT(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// CalcMaxWithdraw is a free data retrieval call binding the contract method 0x4bc9fdc2.
//
// Solidity: function calcMaxWithdraw() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) CalcMaxWithdraw(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "calcMaxWithdraw")
	return *ret0, err
}

// CalcMaxWithdraw is a free data retrieval call binding the contract method 0x4bc9fdc2.
//
// Solidity: function calcMaxWithdraw() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) CalcMaxWithdraw() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.CalcMaxWithdraw(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// CalcMaxWithdraw is a free data retrieval call binding the contract method 0x4bc9fdc2.
//
// Solidity: function calcMaxWithdraw() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) CalcMaxWithdraw() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.CalcMaxWithdraw(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) Confirmations(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "confirmations", arg0, arg1)
	return *ret0, err
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Confirmations(&_MultiSigWalletWithDailyLimit.CallOpts, arg0, arg1)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Confirmations(&_MultiSigWalletWithDailyLimit.CallOpts, arg0, arg1)
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) DailyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "dailyLimit")
	return *ret0, err
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) DailyLimit() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.DailyLimit(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) DailyLimit() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.DailyLimit(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) GetConfirmationCount(opts *bind.CallOpts, transactionId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "getConfirmationCount", transactionId)
	return *ret0, err
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetConfirmationCount(&_MultiSigWalletWithDailyLimit.CallOpts, transactionId)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetConfirmationCount(&_MultiSigWalletWithDailyLimit.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) GetConfirmations(opts *bind.CallOpts, transactionId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "getConfirmations", transactionId)
	return *ret0, err
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetConfirmations(&_MultiSigWalletWithDailyLimit.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetConfirmations(&_MultiSigWalletWithDailyLimit.CallOpts, transactionId)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "getOwners")
	return *ret0, err
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) GetOwners() ([]common.Address, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetOwners(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) GetOwners() ([]common.Address, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetOwners(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) GetTransactionCount(opts *bind.CallOpts, pending bool, executed bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "getTransactionCount", pending, executed)
	return *ret0, err
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetTransactionCount(&_MultiSigWalletWithDailyLimit.CallOpts, pending, executed)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetTransactionCount(&_MultiSigWalletWithDailyLimit.CallOpts, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) GetTransactionIds(opts *bind.CallOpts, from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "getTransactionIds", from, to, pending, executed)
	return *ret0, err
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetTransactionIds(&_MultiSigWalletWithDailyLimit.CallOpts, from, to, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetTransactionIds(&_MultiSigWalletWithDailyLimit.CallOpts, from, to, pending, executed)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) IsConfirmed(opts *bind.CallOpts, transactionId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "isConfirmed", transactionId)
	return *ret0, err
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _MultiSigWalletWithDailyLimit.Contract.IsConfirmed(&_MultiSigWalletWithDailyLimit.CallOpts, transactionId)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _MultiSigWalletWithDailyLimit.Contract.IsConfirmed(&_MultiSigWalletWithDailyLimit.CallOpts, transactionId)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) IsOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "isOwner", arg0)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) IsOwner(arg0 common.Address) (bool, error) {
	return _MultiSigWalletWithDailyLimit.Contract.IsOwner(&_MultiSigWalletWithDailyLimit.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) IsOwner(arg0 common.Address) (bool, error) {
	return _MultiSigWalletWithDailyLimit.Contract.IsOwner(&_MultiSigWalletWithDailyLimit.CallOpts, arg0)
}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) LastDay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "lastDay")
	return *ret0, err
}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) LastDay() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.LastDay(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) LastDay() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.LastDay(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Owners(&_MultiSigWalletWithDailyLimit.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Owners(&_MultiSigWalletWithDailyLimit.CallOpts, arg0)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) Required(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "required")
	return *ret0, err
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) Required() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Required(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) Required() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Required(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) SpentToday(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "spentToday")
	return *ret0, err
}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) SpentToday() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.SpentToday(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) SpentToday() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.SpentToday(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) TransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "transactionCount")
	return *ret0, err
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) TransactionCount() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.TransactionCount(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) TransactionCount() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.TransactionCount(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address destination, uint256 value, bytes data, bool executed)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	ret := new(struct {
		Destination common.Address
		Value       *big.Int
		Data        []byte
		Executed    bool
	})
	out := ret
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, out, "transactions", arg0)
	return *ret, err
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address destination, uint256 value, bytes data, bool executed)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Transactions(&_MultiSigWalletWithDailyLimit.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address destination, uint256 value, bytes data, bool executed)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Transactions(&_MultiSigWalletWithDailyLimit.CallOpts, arg0)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) AddOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "addOwner", owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.AddOwner(&_MultiSigWalletWithDailyLimit.TransactOpts, owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.AddOwner(&_MultiSigWalletWithDailyLimit.TransactOpts, owner)
}

// ChangeDailyLimit is a paid mutator transaction binding the contract method 0xcea08621.
//
// Solidity: function changeDailyLimit(uint256 _dailyLimit) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) ChangeDailyLimit(opts *bind.TransactOpts, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "changeDailyLimit", _dailyLimit)
}

// ChangeDailyLimit is a paid mutator transaction binding the contract method 0xcea08621.
//
// Solidity: function changeDailyLimit(uint256 _dailyLimit) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) ChangeDailyLimit(_dailyLimit *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ChangeDailyLimit(&_MultiSigWalletWithDailyLimit.TransactOpts, _dailyLimit)
}

// ChangeDailyLimit is a paid mutator transaction binding the contract method 0xcea08621.
//
// Solidity: function changeDailyLimit(uint256 _dailyLimit) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) ChangeDailyLimit(_dailyLimit *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ChangeDailyLimit(&_MultiSigWalletWithDailyLimit.TransactOpts, _dailyLimit)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) ChangeRequirement(opts *bind.TransactOpts, _required *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "changeRequirement", _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ChangeRequirement(&_MultiSigWalletWithDailyLimit.TransactOpts, _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ChangeRequirement(&_MultiSigWalletWithDailyLimit.TransactOpts, _required)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) ConfirmTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "confirmTransaction", transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ConfirmTransaction(&_MultiSigWalletWithDailyLimit.TransactOpts, transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ConfirmTransaction(&_MultiSigWalletWithDailyLimit.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) ExecuteTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "executeTransaction", transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ExecuteTransaction(&_MultiSigWalletWithDailyLimit.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ExecuteTransaction(&_MultiSigWalletWithDailyLimit.TransactOpts, transactionId)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) RemoveOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "removeOwner", owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.RemoveOwner(&_MultiSigWalletWithDailyLimit.TransactOpts, owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.RemoveOwner(&_MultiSigWalletWithDailyLimit.TransactOpts, owner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) ReplaceOwner(opts *bind.TransactOpts, owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "replaceOwner", owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ReplaceOwner(&_MultiSigWalletWithDailyLimit.TransactOpts, owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ReplaceOwner(&_MultiSigWalletWithDailyLimit.TransactOpts, owner, newOwner)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) RevokeConfirmation(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "revokeConfirmation", transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.RevokeConfirmation(&_MultiSigWalletWithDailyLimit.TransactOpts, transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.RevokeConfirmation(&_MultiSigWalletWithDailyLimit.TransactOpts, transactionId)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) SubmitTransaction(opts *bind.TransactOpts, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "submitTransaction", destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.SubmitTransaction(&_MultiSigWalletWithDailyLimit.TransactOpts, destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.SubmitTransaction(&_MultiSigWalletWithDailyLimit.TransactOpts, destination, value, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Fallback(&_MultiSigWalletWithDailyLimit.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Fallback(&_MultiSigWalletWithDailyLimit.TransactOpts, calldata)
}

// MultiSigWalletWithDailyLimitConfirmationIterator is returned from FilterConfirmation and is used to iterate over the raw logs and unpacked data for Confirmation events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitConfirmationIterator struct {
	Event *MultiSigWalletWithDailyLimitConfirmation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitConfirmation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitConfirmation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitConfirmation represents a Confirmation event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitConfirmation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfirmation is a free log retrieval operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterConfirmation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MultiSigWalletWithDailyLimitConfirmationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitConfirmationIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "Confirmation", logs: logs, sub: sub}, nil
}

// WatchConfirmation is a free log subscription operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchConfirmation(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitConfirmation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitConfirmation)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Confirmation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfirmation is a log parse operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) ParseConfirmation(log types.Log) (*MultiSigWalletWithDailyLimitConfirmation, error) {
	event := new(MultiSigWalletWithDailyLimitConfirmation)
	if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Confirmation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MultiSigWalletWithDailyLimitDailyLimitChangeIterator is returned from FilterDailyLimitChange and is used to iterate over the raw logs and unpacked data for DailyLimitChange events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitDailyLimitChangeIterator struct {
	Event *MultiSigWalletWithDailyLimitDailyLimitChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitDailyLimitChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitDailyLimitChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitDailyLimitChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitDailyLimitChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitDailyLimitChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitDailyLimitChange represents a DailyLimitChange event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitDailyLimitChange struct {
	DailyLimit *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDailyLimitChange is a free log retrieval operation binding the contract event 0xc71bdc6afaf9b1aa90a7078191d4fc1adf3bf680fca3183697df6b0dc226bca2.
//
// Solidity: event DailyLimitChange(uint256 dailyLimit)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterDailyLimitChange(opts *bind.FilterOpts) (*MultiSigWalletWithDailyLimitDailyLimitChangeIterator, error) {

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "DailyLimitChange")
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitDailyLimitChangeIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "DailyLimitChange", logs: logs, sub: sub}, nil
}

// WatchDailyLimitChange is a free log subscription operation binding the contract event 0xc71bdc6afaf9b1aa90a7078191d4fc1adf3bf680fca3183697df6b0dc226bca2.
//
// Solidity: event DailyLimitChange(uint256 dailyLimit)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchDailyLimitChange(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitDailyLimitChange) (event.Subscription, error) {

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "DailyLimitChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitDailyLimitChange)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "DailyLimitChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDailyLimitChange is a log parse operation binding the contract event 0xc71bdc6afaf9b1aa90a7078191d4fc1adf3bf680fca3183697df6b0dc226bca2.
//
// Solidity: event DailyLimitChange(uint256 dailyLimit)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) ParseDailyLimitChange(log types.Log) (*MultiSigWalletWithDailyLimitDailyLimitChange, error) {
	event := new(MultiSigWalletWithDailyLimitDailyLimitChange)
	if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "DailyLimitChange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MultiSigWalletWithDailyLimitDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitDepositIterator struct {
	Event *MultiSigWalletWithDailyLimitDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitDeposit represents a Deposit event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitDeposit struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*MultiSigWalletWithDailyLimitDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitDepositIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitDeposit)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) ParseDeposit(log types.Log) (*MultiSigWalletWithDailyLimitDeposit, error) {
	event := new(MultiSigWalletWithDailyLimitDeposit)
	if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MultiSigWalletWithDailyLimitExecutionIterator is returned from FilterExecution and is used to iterate over the raw logs and unpacked data for Execution events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitExecutionIterator struct {
	Event *MultiSigWalletWithDailyLimitExecution // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitExecution)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitExecution)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitExecution represents a Execution event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitExecution struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecution is a free log retrieval operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterExecution(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletWithDailyLimitExecutionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitExecutionIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "Execution", logs: logs, sub: sub}, nil
}

// WatchExecution is a free log subscription operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchExecution(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitExecution, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitExecution)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Execution", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecution is a log parse operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) ParseExecution(log types.Log) (*MultiSigWalletWithDailyLimitExecution, error) {
	event := new(MultiSigWalletWithDailyLimitExecution)
	if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Execution", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MultiSigWalletWithDailyLimitExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitExecutionFailureIterator struct {
	Event *MultiSigWalletWithDailyLimitExecutionFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitExecutionFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitExecutionFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitExecutionFailure represents a ExecutionFailure event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitExecutionFailure struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterExecutionFailure(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletWithDailyLimitExecutionFailureIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitExecutionFailureIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitExecutionFailure, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitExecutionFailure)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionFailure is a log parse operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) ParseExecutionFailure(log types.Log) (*MultiSigWalletWithDailyLimitExecutionFailure, error) {
	event := new(MultiSigWalletWithDailyLimitExecutionFailure)
	if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MultiSigWalletWithDailyLimitOwnerAdditionIterator is returned from FilterOwnerAddition and is used to iterate over the raw logs and unpacked data for OwnerAddition events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitOwnerAdditionIterator struct {
	Event *MultiSigWalletWithDailyLimitOwnerAddition // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitOwnerAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitOwnerAddition)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitOwnerAddition)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitOwnerAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitOwnerAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitOwnerAddition represents a OwnerAddition event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitOwnerAddition struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerAddition is a free log retrieval operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterOwnerAddition(opts *bind.FilterOpts, owner []common.Address) (*MultiSigWalletWithDailyLimitOwnerAdditionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitOwnerAdditionIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "OwnerAddition", logs: logs, sub: sub}, nil
}

// WatchOwnerAddition is a free log subscription operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchOwnerAddition(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitOwnerAddition, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitOwnerAddition)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnerAddition is a log parse operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) ParseOwnerAddition(log types.Log) (*MultiSigWalletWithDailyLimitOwnerAddition, error) {
	event := new(MultiSigWalletWithDailyLimitOwnerAddition)
	if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MultiSigWalletWithDailyLimitOwnerRemovalIterator is returned from FilterOwnerRemoval and is used to iterate over the raw logs and unpacked data for OwnerRemoval events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitOwnerRemovalIterator struct {
	Event *MultiSigWalletWithDailyLimitOwnerRemoval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitOwnerRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitOwnerRemoval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitOwnerRemoval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitOwnerRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitOwnerRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitOwnerRemoval represents a OwnerRemoval event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitOwnerRemoval struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerRemoval is a free log retrieval operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterOwnerRemoval(opts *bind.FilterOpts, owner []common.Address) (*MultiSigWalletWithDailyLimitOwnerRemovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitOwnerRemovalIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "OwnerRemoval", logs: logs, sub: sub}, nil
}

// WatchOwnerRemoval is a free log subscription operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchOwnerRemoval(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitOwnerRemoval, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitOwnerRemoval)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnerRemoval is a log parse operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) ParseOwnerRemoval(log types.Log) (*MultiSigWalletWithDailyLimitOwnerRemoval, error) {
	event := new(MultiSigWalletWithDailyLimitOwnerRemoval)
	if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MultiSigWalletWithDailyLimitRequirementChangeIterator is returned from FilterRequirementChange and is used to iterate over the raw logs and unpacked data for RequirementChange events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitRequirementChangeIterator struct {
	Event *MultiSigWalletWithDailyLimitRequirementChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitRequirementChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitRequirementChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitRequirementChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitRequirementChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitRequirementChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitRequirementChange represents a RequirementChange event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitRequirementChange struct {
	Required *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequirementChange is a free log retrieval operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterRequirementChange(opts *bind.FilterOpts) (*MultiSigWalletWithDailyLimitRequirementChangeIterator, error) {

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitRequirementChangeIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "RequirementChange", logs: logs, sub: sub}, nil
}

// WatchRequirementChange is a free log subscription operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchRequirementChange(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitRequirementChange) (event.Subscription, error) {

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitRequirementChange)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "RequirementChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequirementChange is a log parse operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) ParseRequirementChange(log types.Log) (*MultiSigWalletWithDailyLimitRequirementChange, error) {
	event := new(MultiSigWalletWithDailyLimitRequirementChange)
	if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "RequirementChange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MultiSigWalletWithDailyLimitRevocationIterator is returned from FilterRevocation and is used to iterate over the raw logs and unpacked data for Revocation events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitRevocationIterator struct {
	Event *MultiSigWalletWithDailyLimitRevocation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitRevocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitRevocation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitRevocation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitRevocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitRevocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitRevocation represents a Revocation event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitRevocation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRevocation is a free log retrieval operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterRevocation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MultiSigWalletWithDailyLimitRevocationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitRevocationIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "Revocation", logs: logs, sub: sub}, nil
}

// WatchRevocation is a free log subscription operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchRevocation(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitRevocation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitRevocation)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Revocation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevocation is a log parse operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) ParseRevocation(log types.Log) (*MultiSigWalletWithDailyLimitRevocation, error) {
	event := new(MultiSigWalletWithDailyLimitRevocation)
	if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Revocation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MultiSigWalletWithDailyLimitSubmissionIterator is returned from FilterSubmission and is used to iterate over the raw logs and unpacked data for Submission events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitSubmissionIterator struct {
	Event *MultiSigWalletWithDailyLimitSubmission // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitSubmission)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitSubmission)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitSubmission represents a Submission event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitSubmission struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSubmission is a free log retrieval operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterSubmission(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletWithDailyLimitSubmissionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitSubmissionIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "Submission", logs: logs, sub: sub}, nil
}

// WatchSubmission is a free log subscription operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchSubmission(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitSubmission, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitSubmission)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Submission", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubmission is a log parse operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) ParseSubmission(log types.Log) (*MultiSigWalletWithDailyLimitSubmission, error) {
	event := new(MultiSigWalletWithDailyLimitSubmission)
	if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Submission", log); err != nil {
		return nil, err
	}
	return event, nil
}
