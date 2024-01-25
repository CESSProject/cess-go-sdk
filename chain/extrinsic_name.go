/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

var ExtrinsicsName map[types.CallIndex]string

const (
	// Assets
	ExtName_Assets_approve_transfer      = "Assets.approve_transfer"
	ExtName_Assets_block                 = "Assets.block"
	ExtName_Assets_burn                  = "Assets.burn"
	ExtName_Assets_cancel_approval       = "Assets.cancel_approval"
	ExtName_Assets_clear_metadata        = "Assets.clear_metadata"
	ExtName_Assets_create                = "Assets.create"
	ExtName_Assets_destroy_accounts      = "Assets.destroy_accounts"
	ExtName_Assets_destroy_approvals     = "Assets.destroy_approvals"
	ExtName_Assets_finish_destroy        = "Assets.finish_destroy"
	ExtName_Assets_force_asset_status    = "Assets.force_asset_status"
	ExtName_Assets_force_cancel_approval = "Assets.force_cancel_approval"
	ExtName_Assets_force_clear_metadata  = "Assets.force_clear_metadata"
	ExtName_Assets_force_create          = "Assets.force_create"
	ExtName_Assets_force_set_metadata    = "Assets.force_set_metadata"
	ExtName_Assets_force_transfer        = "Assets.force_transfer"
	ExtName_Assets_freeze                = "Assets.freeze"
	ExtName_Assets_freeze_asset          = "Assets.freeze_asset"
	ExtName_Assets_mint                  = "Assets.mint"
	ExtName_Assets_refund                = "Assets.refund"
	ExtName_Assets_refund_other          = "Assets.refund_other"
	ExtName_Assets_set_metadata          = "Assets.set_metadata"
	ExtName_Assets_set_min_balance       = "Assets.set_min_balance"
	ExtName_Assets_set_team              = "Assets.set_team"
	ExtName_Assets_start_destroy         = "Assets.start_destroy"
	ExtName_Assets_thaw                  = "Assets.thaw"
	ExtName_Assets_thaw_asset            = "Assets.thaw_asset"
	ExtName_Assets_touch                 = "Assets.touch"
	ExtName_Assets_touch_other           = "Assets.touch_other"
	ExtName_Assets_transfer              = "Assets.transfer"
	ExtName_Assets_transfer_approved     = "Assets.transfer_approved"
	ExtName_Assets_transfer_keep_alive   = "Assets.transfer_keep_alive"

	// Audit
	ExtName_Audit_submit_idle_proof            = "Audit.submit_idle_proof"
	ExtName_Audit_submit_service_proof         = "Audit.submit_service_proof"
	ExtName_Audit_submit_verify_idle_result    = "Audit.submit_verify_idle_result"
	ExtName_Audit_submit_verify_service_result = "Audit.submit_verify_service_result"
	ExtName_Audit_test_update_clear_slip       = "Audit.test_update_clear_slip"
	ExtName_Audit_test_update_verify_slip      = "Audit.test_update_verify_slip"
	ExtName_Audit_update_counted_clear         = "Audit.update_counted_clear"

	// Babe
	ExtName_Babe_plan_config_change           = "Babe.plan_config_change"
	ExtName_Babe_report_equivocation          = "Babe.report_equivocation"
	ExtName_Babe_report_equivocation_unsigned = "Babe.report_equivocation_unsigned"

	// Balances
	ExtName_Balances_force_set_balance      = "Balances.force_set_balance"
	ExtName_Balances_force_transfer         = "Balances.force_transfer"
	ExtName_Balances_force_unreserve        = "Balances.force_unreserve"
	ExtName_Balances_set_balance_deprecated = "Balances.set_balance_deprecated"
	ExtName_Balances_transfer               = "Balances.transfer"
	ExtName_Balances_transfer_all           = "Balances.transfer_all"
	ExtName_Balances_transfer_allow_death   = "Balances.transfer_allow_death"
	ExtName_Balances_transferKeepAlive      = "Balances.transfer_keep_alive"
	ExtName_Balances_upgrade_accounts       = "Balances.upgrade_accounts"

	// BaseFee
	ExtName_BaseFee_set_base_fee_per_gas = "BaseFee.set_base_fee_per_gas"
	ExtName_BaseFee_set_elasticity       = "BaseFee.set_elasticity"

	// Bounties
	ExtName_Bounties_accept_curator       = "Bounties.accept_curator"
	ExtName_Bounties_approve_bounty       = "Bounties.approve_bounty"
	ExtName_Bounties_award_bounty         = "Bounties.award_bounty"
	ExtName_Bounties_claim_bounty         = "Bounties.claim_bounty"
	ExtName_Bounties_close_bounty         = "Bounties.close_bounty"
	ExtName_Bounties_extend_bounty_expiry = "Bounties.extend_bounty_expiry"
	ExtName_Bounties_propose_bounty       = "Bounties.propose_bounty"
	ExtName_Bounties_propose_curator      = "Bounties.propose_curator"
	ExtName_Bounties_unassign_curator     = "Bounties.unassign_curator"

	// Cacher
	ExtName_Cacher_logout   = "Cacher.logout"
	ExtName_Cacher_pay      = "Cacher.pay"
	ExtName_Cacher_register = "Cacher.register"
	ExtName_Cacher_update   = "Cacher.update"

	// CesMq
	ExtName_CesMq_force_push_pallet_message = "CesMq.force_push_pallet_message"
	ExtName_CesMq_push_message              = "CesMq.push_message"
	ExtName_CesMq_sync_offchain_message     = "CesMq.sync_offchain_message"

	// CessTreasury
	ExtName_CessTreasury_pid_burn_funds    = "CessTreasury.pid_burn_funds"
	ExtName_CessTreasury_pid_send_funds    = "CessTreasury.pid_send_funds"
	ExtName_CessTreasury_send_funds_to_pid = "CessTreasury.send_funds_to_pid"
	ExtName_CessTreasury_send_funds_to_sid = "CessTreasury.send_funds_to_sid"
	ExtName_CessTreasury_sid_burn_funds    = "CessTreasury.sid_burn_funds"
	ExtName_CessTreasury_sid_send_funds    = "CessTreasury.sid_send_funds"

	// ChildBounties
	ExtName_ChildBounties_accept_curator     = "ChildBounties.accept_curator"
	ExtName_ChildBounties_add_child_bounty   = "ChildBounties.add_child_bounty"
	ExtName_ChildBounties_award_child_bounty = "ChildBounties.award_child_bounty"
	ExtName_ChildBounties_claim_child_bounty = "ChildBounties.claim_child_bounty"
	ExtName_ChildBounties_close_child_bounty = "ChildBounties.close_child_bounty"
	ExtName_ChildBounties_propose_curator    = "ChildBounties.propose_curator"
	ExtName_ChildBounties_unassign_curator   = "ChildBounties.unassign_curator"

	// Contracts
	ExtName_Contracts_call                             = "Contracts.call"
	ExtName_Contracts_call_old_weight                  = "Contracts.call_old_weight"
	ExtName_Contracts_instantiate                      = "Contracts.instantiate"
	ExtName_Contracts_instantiate_old_weight           = "Contracts.instantiate_old_weight"
	ExtName_Contracts_instantiate_with_code            = "Contracts.instantiate_with_code"
	ExtName_Contracts_instantiate_with_code_old_weight = "Contracts.instantiate_with_code_old_weight"
	ExtName_Contracts_migrate                          = "Contracts.migrate"
	ExtName_Contracts_remove_code                      = "Contracts.remove_code"
	ExtName_Contracts_set_code                         = "Contracts.set_code"
	ExtName_Contracts_upload_code                      = "Contracts.upload_code"

	// Council
	ExtName_Council_close               = "Council.close"
	ExtName_Council_disapprove_proposal = "Council.disapprove_proposal"
	ExtName_Council_execute             = "Council.execute"
	ExtName_Council_propose             = "Council.propose"
	ExtName_Council_set_members         = "Council.set_members"
	ExtName_Council_vote                = "Council.vote"

	// DynamicFee
	ExtName_DynamicFee_min_gas_price        = "DynamicFee.min_gas_price"
	ExtName_DynamicFee_pallet_version       = "DynamicFee.pallet_version"
	ExtName_DynamicFee_target_min_gas_price = "DynamicFee.target_min_gas_price"

	// ElectionProviderMultiPhase
	ExtName_ElectionProviderMultiPhase_current_phase                = "ElectionProviderMultiPhase.current_phase"
	ExtName_ElectionProviderMultiPhase_desired_targets              = "ElectionProviderMultiPhase.desired_targets"
	ExtName_ElectionProviderMultiPhase_minimum_untrusted_score      = "ElectionProviderMultiPhase.minimum_untrusted_score"
	ExtName_ElectionProviderMultiPhase_pallet_version               = "ElectionProviderMultiPhase.pallet_version"
	ExtName_ElectionProviderMultiPhase_queued_solution              = "ElectionProviderMultiPhase.queued_solution"
	ExtName_ElectionProviderMultiPhase_round                        = "ElectionProviderMultiPhase.round"
	ExtName_ElectionProviderMultiPhase_signed_submission_indices    = "ElectionProviderMultiPhase.signed_submission_indices"
	ExtName_ElectionProviderMultiPhase_signed_submission_next_index = "ElectionProviderMultiPhase.signed_submission_next_index"
	ExtName_ElectionProviderMultiPhase_signed_submissions_map       = "ElectionProviderMultiPhase.signed_submissions_map"
	ExtName_ElectionProviderMultiPhase_snapshot                     = "ElectionProviderMultiPhase.snapshot"
	ExtName_ElectionProviderMultiPhase_snapshot_metadata            = "ElectionProviderMultiPhase.snapshot_metadata"

	// Ethereum
	ExtName_Ethereum_blockHash                    = "Ethereum.block_hash"
	ExtName_Ethereum_current_block                = "Ethereum.current_block"
	ExtName_Ethereum_current_receipts             = "Ethereum.current_receipts"
	ExtName_Ethereum_current_transaction_statuses = "Ethereum.current_transaction_statuses"
	ExtName_Ethereum_pallet_version               = "Ethereum.pallet_version"
	ExtName_Ethereum_pending                      = "Ethereum.pending"

	// Evm
	ExtName_Evm_account_codes          = "Evm.account_codes"
	ExtName_Evm_account_codes_metadata = "Evm.account_codes_metadata"
	ExtName_Evm_account_storages       = "Evm.account_storages"
	ExtName_Evm_pallet_version         = "Evm.pallet_version"
	ExtName_Evm_suicided               = "Evm.suicided"

	// EvmAccountMapping
	ExtName_EvmAccountMapping_account_nonce  = "EvmAccountMapping.account_nonce"
	ExtName_EvmAccountMapping_pallet_version = "EvmAccountMapping.pallet_version"

	// EvmChainId
	ExtName_EvmChainId_chain_id       = "EvmChainId.chain_id"
	ExtName_EvmChainId_pallet_version = "EvmChainId.pallet_version"

	// FileBank
	ExtName_FileBank_calculate_report             = "FileBank.calculate_report"
	ExtName_FileBank_cert_idle_space              = "FileBank.cert_idle_space"
	ExtName_FileBank_claim_restoral_noexist_order = "FileBank.claim_restoral_noexist_order"
	ExtName_FileBank_claim_restoral_order         = "FileBank.claim_restoral_order"
	ExtName_FileBank_create_bucket                = "FileBank.create_bucket"
	ExtName_FileBank_delete_bucket                = "FileBank.delete_bucket"
	ExtName_FileBank_delete_file                  = "FileBank.delete_file"
	ExtName_FileBank_generate_restoral_order      = "FileBank.generate_restoral_order"
	ExtName_FileBank_ownership_transfer           = "FileBank.ownership_transfer"
	ExtName_FileBank_replace_idle_space           = "FileBank.replace_idle_space"
	ExtName_FileBank_restoral_order_complete      = "FileBank.restoral_order_complete"
	ExtName_FileBank_root_clear_file              = "FileBank.root_clear_file"
	ExtName_FileBank_transfer_report              = "FileBank.transfer_report"
	ExtName_FileBank_upload_declaration           = "FileBank.upload_declaration"

	// Grandpa
	ExtName_Grandpa_note_stalled                 = "Grandpa.note_stalled"
	ExtName_Grandpa_report_equivocation          = "Grandpa.report_equivocation"
	ExtName_Grandpa_report_equivocation_unsigned = "Grandpa.report_equivocation_unsigned"

	// ImOnline
	ExtName_ImOnline_heartbeat = "ImOnline.heartbeat"

	// Indices
	ExtName_Indices_claim          = "Indices.claim"
	ExtName_Indices_force_transfer = "Indices.force_transfer"
	ExtName_Indices_free           = "Indices.free"
	ExtName_Indices_freeze         = "Indices.freeze"
	ExtName_Indices_transfer       = "Indices.transfer"

	// Multisig
	ExtName_Multisig_approve_as_multi    = "Multisig.approve_as_multi"
	ExtName_Multisig_as_multi            = "Multisig.as_multi"
	ExtName_Multisig_as_multi_threshold1 = "Multisig.as_multi_threshold1"
	ExtName_Multisig_cancel_as_multi     = "Multisig.cancel_as_multi"

	// Oss
	ExtName_Oss_authorize        = "Oss.authorize"
	ExtName_Oss_cancel_authorize = "Oss.cancel_authorize"
	ExtName_Oss_destroy          = "Oss.destroy"
	ExtName_Oss_register         = "Oss.register"
	ExtName_Oss_update           = "Oss.update"

	// Preimage
	ExtName_Preimage_note_preimage      = "Preimage.note_preimage"
	ExtName_Preimage_request_preimage   = "Preimage.request_preimage"
	ExtName_Preimage_unnote_preimage    = "Preimage.unnote_preimage"
	ExtName_Preimage_unrequest_preimage = "Preimage.unrequest_preimage"

	// Proxy
	ExtName_Proxy_add_proxy           = "Proxy.add_proxy"
	ExtName_Proxy_announce            = "Proxy.announce"
	ExtName_Proxy_create_pure         = "Proxy.create_pure"
	ExtName_Proxy_kill_pure           = "Proxy.kill_pure"
	ExtName_Proxy_proxy               = "Proxy.proxy"
	ExtName_Proxy_proxy_announced     = "Proxy.proxy_announced"
	ExtName_Proxy_reject_announcement = "Proxy.reject_announcement"
	ExtName_Proxy_remove_announcement = "Proxy.remove_announcement"
	ExtName_Proxy_remove_proxies      = "Proxy.remove_proxies"
	ExtName_Proxy_remove_proxy        = "Proxy.remove_proxy"

	// Scheduler
	ExtName_Scheduler_cancel               = "Scheduler.cancel"
	ExtName_Scheduler_cancel_named         = "Scheduler.cancel_named"
	ExtName_Scheduler_schedule             = "Scheduler.schedule"
	ExtName_Scheduler_schedule_after       = "Scheduler.schedule_after"
	ExtName_Scheduler_schedule_named       = "Scheduler.schedule_named"
	ExtName_Scheduler_schedule_named_after = "Scheduler.schedule_named_after"

	// Session
	ExtName_Session_purge_keys = "Session.purge_keys"
	ExtName_Session_set_keys   = "Session.set_keys"

	// Sminer
	ExtName_Sminer_clear_miner_service        = "Sminer.clear_miner_service"
	ExtName_Sminer_faucet                     = "Sminer.faucet"
	ExtName_Sminer_faucet_top_up              = "Sminer.faucet_top_up"
	ExtName_Sminer_increase_collateral        = "Sminer.increase_collateral"
	ExtName_Sminer_increase_declaration_space = "Sminer.increase_declaration_space"
	ExtName_Sminer_miner_exit                 = "Sminer.miner_exit"
	ExtName_Sminer_miner_exit_prep            = "Sminer.miner_exit_prep"
	ExtName_Sminer_miner_withdraw             = "Sminer.miner_withdraw"
	ExtName_Sminer_receive_reward             = "Sminer.receive_reward"
	ExtName_Sminer_register_pois_key          = "Sminer.register_pois_key"
	ExtName_Sminer_regnstk                    = "Sminer.regnstk"
	ExtName_Sminer_regnstk_assign_staking     = "Sminer.regnstk_assign_staking"
	ExtName_Sminer_update_beneficiary         = "Sminer.update_beneficiary"
	ExtName_Sminer_update_expender            = "Sminer.update_expender"
	ExtName_Sminer_update_peer_id             = "Sminer.update_peer_id"

	// Staking
	ExtName_Staking_bond                  = "Staking.bond"
	ExtName_Staking_bond_extra            = "Staking.bond_extra"
	ExtName_Staking_cancel_deferred_slash = "Staking.cancel_deferred_slash"
	ExtName_Staking_chill                 = "Staking.chill"
	ExtName_Staking_chill_other           = "Staking.chill_other"
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."
	// ExtName_Staking_                      = "Staking."

	// Timestamp
	ExtName_Timestamp_set = "Timestamp.set"
)

func (c *chainClient) InitExtrinsicsName() error {
	ExtrinsicsName = make(map[types.CallIndex]string, 0)
	// Au
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Timestamp_set); err == nil {
		//fmt.Println(ExtName_Timestamp_set, ".callIndex.MethodIndex:", callIndex.MethodIndex)
		//fmt.Println(ExtName_Timestamp_set, ".callIndex.SectionIndex:", callIndex.SectionIndex)
		ExtrinsicsName[callIndex] = ExtName_Timestamp_set
	} else {
		return err
	}

	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Balances_transferKeepAlive); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Balances_transferKeepAlive
	} else {
		return err
	}

	// BaseFee
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_BaseFee_set_base_fee_per_gas); err == nil {
		ExtrinsicsName[callIndex] = ExtName_BaseFee_set_base_fee_per_gas
	} else {
		return err
	}

	return nil
}
