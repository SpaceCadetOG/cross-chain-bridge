// SPDX-License-Identifier: BicBlockchainSolutions
pragma solidity 0.8.10;

/*
 * IERC20
 */
function transferFrom(address from, address to,uint256 amount) external returns (bool);

contract HTLC {
  uint256 public startTime;
  uint256 public lockTime = 10000 seconds;
  string public secret;
  bytes32 public hash;
  address public reciever;
  address public owner;
  uint256 public amount;

  // IERC20

  constructor(address _reciever, uint256 _amount) {
    reciever = _reciever;
    owner = msg.sender;
    amount = _amount;
    // token
  }

  function send() external {}

  function withdraw() external {}

  function refund() external {}
}
