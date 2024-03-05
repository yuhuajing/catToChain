// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract carTrace {
    struct Car {
        string[] prodinfo;
        string[] logisinfo;
        string[] saleinfo;
        string[] repairinfo;
    }
    address public owner;
    address public prodmanager;
    address public logismanager;
    address public salemanager;
    address public repairmanager;

    mapping(string => Car) carTraceInfo;

    constructor(
        address _prodmanager,
        address _logismanager,
        address _salemanager,
        address _repairmanager
    ) {
        owner = msg.sender;
        prodmanager = _prodmanager;
        logismanager = _logismanager;
        salemanager = _salemanager;
        repairmanager = _repairmanager;
    }

    modifier onlyProdManager() {
        require(msg.sender == prodmanager, "ONLY_PROD_MANAGER_ROLE");
        _;
    }
    modifier onlyOwner() {
        require(msg.sender == owner, "ONLY_OWNER_ROLE");
        _;
    }
    modifier onlyLogisManager() {
        require(msg.sender == logismanager, "ONLY_LOGIS_MANAGER_ROLE");
        _;
    }
    modifier onlySaleManager() {
        require(msg.sender == salemanager, "ONLY_SALE_MANAGER_ROLE");
        _;
    }
    modifier onlyRepaieManager() {
        require(msg.sender == repairmanager, "ONLY_REPAIR_MANAGER_ROLE");
        _;
    }

    function setProdManager(address _r) external onlyOwner {
        prodmanager = _r;
    }

    function setLogisManager(address _r) external onlyOwner {
        logismanager = _r;
    }

    function setSaleManager(address _r) external onlyOwner {
        salemanager = _r;
    }

    function setRepairManager(address _r) external onlyOwner {
        repairmanager = _r;
    }

    receive() external payable {}

    function addProdData(string memory carId, string memory description)
    external
    onlyProdManager
    {
        carTraceInfo[carId].prodinfo.push(description);
    }

    function getProdData(string memory carId)
    external
    view
    returns (string[] memory)
    {
        return carTraceInfo[carId].prodinfo;
    }

    function addLogisData(string memory carId, string memory description)
    external
    onlyLogisManager
    {
        carTraceInfo[carId].logisinfo.push(description);
    }

    function getLogisData(string memory carId)
    external
    view
    returns (string[] memory)
    {
        return carTraceInfo[carId].logisinfo;
    }

    function addSaleData(string memory carId, string memory description)
    external
    onlySaleManager
    {
        carTraceInfo[carId].saleinfo.push(description);
    }

    function getSaleData(string memory carId)
    external
    view
    returns (string[] memory)
    {
        return carTraceInfo[carId].saleinfo;
    }

    function addRepairData(string memory carId, string memory description)
    external
    onlyRepaieManager
    {
        carTraceInfo[carId].repairinfo.push(description);
    }

    function getRepairData(string memory carId)
    external
    view
    returns (string[] memory)
    {
        return carTraceInfo[carId].repairinfo;
    }
}
