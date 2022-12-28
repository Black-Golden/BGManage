package db

import (
	"fmt"
	"github.com/ethereum/coin-manage/types"
)

func (m *Mysql) QueryCoinHolderCount(contractAddr string) (int, error) {
	count := 0
	sql := fmt.Sprintf("select count(*) from balance_erc20 where contract_addr = \"%s\"", contractAddr)
	ok, err := m.engine.SQL(sql).Limit(1).Get(&count)
	if err != nil {
		return count, err
	}
	if !ok {
		return count, nil
	}

	return count, err
}

func (m *Mysql) QueryCoinHolders(contractAddr string) ([]*types.Balance_Erc20, error) {
	balances := make([]*types.Balance_Erc20, 0)
	err := m.engine.Table("balance_erc20").Where("contract_addr = ?", contractAddr).Find(&balances)
	if err != nil {
		return nil, err
	}
	return balances, err
}

func (m *Mysql) QueryTxHistory(accountAddr string) ([]*types.Tx, error) {
	txs := make([]*types.Tx, 0)
	err := m.engine.Where("addr_from = ? or addr_to = ? ", accountAddr, accountAddr).Find(&txs)
	if err != nil {
		return nil, err
	}
	return txs, err
}

func (m *Mysql) QueryTxErc20History(accountAddr string) ([]*types.Erc20Tx, error) {
	txs := make([]*types.Erc20Tx, 0)
	err := m.engine.Table("tx_erc20").Where("sender = ? or receiver = ? ", accountAddr, accountAddr).Find(&txs)
	if err != nil {
		return nil, err
	}
	return txs, err
}

func (m *Mysql) QuerySpecifyCoinInfo(contractAddr string) (*types.Erc20Info, error) {
	info := &types.Erc20Info{}
	ok, err := m.engine.Table("erc20_info").Where("addr = ?", contractAddr).Get(info)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	return info, err
}

func (m *Mysql) QueryCoinInfos(accountAddr string) ([]*types.Erc20Info, error) {
	infos := make([]*types.Erc20Info, 0)
	err := m.engine.Join("INNER", "balance_erc20", "balance_erc20.contract_addr = erc20_info.addr").Where("balance_erc20.addr = ? ", accountAddr).Find(&infos)
	if err != nil {
		return nil, err
	}
	return infos, err
}

func (m *Mysql) QueryAllCoinAllHolders(accountAddr string) (int, error) {
	count := 0
	str := fmt.Sprintf("SELECT count(*) FROM balance_erc20 where contract_addr  in (SELECT contract_addr FROM balance_erc20 where addr = \"%s\")", accountAddr)
	ok, err := m.engine.SQL(str).Get(&count)
	if err != nil {
		return -1, err
	}
	if !ok {
		return -1, nil
	}
	return count, err
}

func (m *Mysql) QueryABI(contractAddr string) (*types.ContractAbi, error) {
	var abi *types.ContractAbi
	err := m.engine.Table("contract_abi").Where("contract_addr = ? ", contractAddr).Find(&abi)
	if err != nil {
		return nil, err
	}
	return abi, err
}

func (m *Mysql) QueryReceiver(contractAddr string) (types.ContractReceiver, error) {
	var receiver types.ContractReceiver
	sql := fmt.Sprintf("select * from t_receiver where contract_addr = \"%s\"", contractAddr)
	ok, err := m.engine.SQL(sql).Limit(1).Get(&receiver)
	if err != nil {
		return receiver, err
	}
	if !ok {
		return receiver, nil
	}
	return receiver, nil
}