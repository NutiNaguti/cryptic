// SPDX-License-Identifier: MIT
pragma solidity ^0.8.1;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract CrypticERC20 is ERC20, Ownable {
  event Minted(address to, uint256 amount);

  constructor() ERC20("Cryptic", "CPT") {}

  function mint(address to, uint256 amount) public onlyOwner {
    _mint(to, amount);
    emit Minted(to, amount);
  }
}
