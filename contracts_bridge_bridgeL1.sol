pragma solidity ^0.8.0;

contract BridgeL1 {
    event Deposit(address indexed user, uint256 amount);
    event Withdraw(address indexed user, uint256 amount);

    function deposit() external payable {
        emit Deposit(msg.sender, msg.value);
    }

    function withdraw(uint256 amount) external {
        payable(msg.sender).transfer(amount);
        emit Withdraw(msg.sender, amount);
    }
}
