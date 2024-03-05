// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract carTrace {
    struct Car {
        string[] prodinfo;
        string[] logisinfo;
        string[] saleinfo;
        string[] repairinfo;
    }

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
        prodmanager = _prodmanager;
        logismanager = _logismanager;
        salemanager = _salemanager;
        repairmanager = _repairmanager;
    }

    modifier onlyProdManager() {
        require(msg.sender == prodmanager, "ONLY_PROD_MANAGER_ROLE");
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

    receive() external payable {}

    function addProdData(string memory carId, string memory description)
    onlyProdManager external
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
    onlyLogisManager external
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
    onlySaleManager external
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
    onlyRepaieManager external
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
