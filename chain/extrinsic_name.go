/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"github.com/AstaFrode/go-substrate-rpc-client/v4/types"
)

var ExtrinsicsName map[types.CallIndex]string

const (
	// AssetConversion
	ExtName_AssetConversion_add_liquidity                = "AssetConversion.add_liquidity"
	ExtName_AssetConversion_create_pool                  = "AssetConversion.create_pool"
	ExtName_AssetConversion_remove_liquidity             = "AssetConversion.remove_liquidity"
	ExtName_AssetConversion_swap_exact_tokens_for_tokens = "AssetConversion.swap_exact_tokens_for_tokens"
	ExtName_AssetConversion_swap_tokens_for_exact_tokens = "AssetConversion.swap_tokens_for_exact_tokens"
	ExtName_AssetConversion_stouch                       = "AssetConversion.touch"

	// AssetRate
	ExtName_AssetRate_create = "AssetRate.create"
	ExtName_AssetRate_remove = "AssetRate.remove"
	ExtName_AssetRate_update = "AssetRate.update"

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
	ExtName_Assets_transfer_all          = "Assets.transfer_all"
	ExtName_Assets_transfer_approved     = "Assets.transfer_approved"
	ExtName_Assets_transfer_keep_alive   = "Assets.transfer_keep_alive"
	ExtName_Assets_transfer_ownership    = "Assets.transfer_ownership"

	// Audit
	ExtName_Audit_point_miner_challenge        = "Audit.point_miner_challenge"
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
	ExtName_Balances_burn                        = "Balances.burn"
	ExtName_Balances_force_adjust_total_issuance = "Balances.force_adjust_total_issuance"
	ExtName_Balances_force_set_balance           = "Balances.force_set_balance"
	ExtName_Balances_force_transfer              = "Balances.force_transfer"
	ExtName_Balances_force_unreserve             = "Balances.force_unreserve"
	ExtName_Balances_transfer_all                = "Balances.transfer_all"
	ExtName_Balances_transfer_allow_death        = "Balances.transfer_allow_death"
	ExtName_Balances_transferKeepAlive           = "Balances.transfer_keep_alive"
	ExtName_Balances_upgrade_accounts            = "Balances.upgrade_accounts"

	// BaseFee
	ExtName_BaseFee_set_base_fee_per_gas = "BaseFee.set_base_fee_per_gas"
	ExtName_BaseFee_set_elasticity       = "BaseFee.set_elasticity"

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

	// ElectionProviderMultiPhase
	ExtName_ElectionProviderMultiPhase_governance_fallback           = "ElectionProviderMultiPhase.governance_fallback"
	ExtName_ElectionProviderMultiPhase_set_emergency_election_result = "ElectionProviderMultiPhase.set_emergency_election_result"
	ExtName_ElectionProviderMultiPhase_set_minimum_untrusted_score   = "ElectionProviderMultiPhase.set_minimum_untrusted_score"
	ExtName_ElectionProviderMultiPhase_submit                        = "ElectionProviderMultiPhase.submit"
	ExtName_ElectionProviderMultiPhase_submit_unsigned               = "ElectionProviderMultiPhase.submit_unsigned"

	// Ethereum
	ExtName_Ethereum_transact = "Ethereum.transact"

	// EVM
	ExtName_Evm_call     = "EVM.call"
	ExtName_Evm_create   = "EVM.create"
	ExtName_Evm_create2  = "EVM.create2"
	ExtName_Evm_withdraw = "EVM.withdraw"

	// EvmAccountMapping
	ExtName_EvmAccountMapping_meta_call = "EvmAccountMapping.meta_call"

	// FastUnstake
	ExtName_FastUnstake_control               = "FastUnstake.control"
	ExtName_FastUnstake_deregister            = "FastUnstake.deregister"
	ExtName_FastUnstake_register_fast_unstake = "FastUnstake.register_fast_unstake"

	// FileBank
	ExtName_FileBank_calculate_report             = "FileBank.calculate_report"
	ExtName_FileBank_cert_idle_space              = "FileBank.cert_idle_space"
	ExtName_FileBank_claim_restoral_noexist_order = "FileBank.claim_restoral_noexist_order"
	ExtName_FileBank_claim_restoral_order         = "FileBank.claim_restoral_order"
	ExtName_FileBank_delete_file                  = "FileBank.delete_file"
	ExtName_FileBank_generate_restoral_order      = "FileBank.generate_restoral_order"
	ExtName_FileBank_replace_idle_space           = "FileBank.replace_idle_space"
	ExtName_FileBank_restoral_order_complete      = "FileBank.restoral_order_complete"
	ExtName_FileBank_root_clear_file              = "FileBank.root_clear_file"
	ExtName_FileBank_territory_file_delivery      = "FileBank.territory_file_delivery"
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

	// MultiBlockMigrations
	ExtName_MultiBlockMigrations_clear_historic          = "MultiBlockMigrations.clear_historic"
	ExtName_MultiBlockMigrations_force_onboard_mbms      = "MultiBlockMigrations.force_onboard_mbms"
	ExtName_MultiBlockMigrations_force_set_active_cursor = "MultiBlockMigrations.force_set_active_cursor"
	ExtName_MultiBlockMigrations_force_set_cursor        = "MultiBlockMigrations.force_set_cursor"

	// Multisig
	ExtName_Multisig_approve_as_multi    = "Multisig.approve_as_multi"
	ExtName_Multisig_as_multi            = "Multisig.as_multi"
	ExtName_Multisig_as_multi_threshold1 = "Multisig.as_multi_threshold_1"
	ExtName_Multisig_cancel_as_multi     = "Multisig.cancel_as_multi"

	// Oss
	ExtName_Oss_authorize           = "Oss.authorize"
	ExtName_Oss_cancel_authorize    = "Oss.cancel_authorize"
	ExtName_Oss_destroy             = "Oss.destroy"
	ExtName_Oss_evm_proxy_authorzie = "Oss.evm_proxy_authorzie"
	ExtName_Oss_proxy_authorzie     = "Oss.proxy_authorzie"
	ExtName_Oss_register            = "Oss.register"
	ExtName_Oss_update              = "Oss.update"

	// Parameters
	ExtName_Parameters_set_parameter = "Parameters.set_parameter"

	// PoolAssets
	ExtName_PoolAssets_approve_transfer      = "PoolAssets.approve_transfer"
	ExtName_PoolAssets_block                 = "PoolAssets.block"
	ExtName_PoolAssets_burn                  = "PoolAssets.burn"
	ExtName_PoolAssets_cancel_approval       = "PoolAssets.cancel_approval"
	ExtName_PoolAssets_clear_metadata        = "PoolAssets.clear_metadata"
	ExtName_PoolAssets_create                = "PoolAssets.create"
	ExtName_PoolAssets_destroy_accounts      = "PoolAssets.destroy_accounts"
	ExtName_PoolAssets_destroy_approvals     = "PoolAssets.destroy_approvals"
	ExtName_PoolAssets_finish_destroy        = "PoolAssets.finish_destroy"
	ExtName_PoolAssets_force_asset_status    = "PoolAssets.force_asset_status"
	ExtName_PoolAssets_force_cancel_approval = "PoolAssets.force_cancel_approval"
	ExtName_PoolAssets_force_clear_metadata  = "PoolAssets.force_clear_metadata"
	ExtName_PoolAssets_force_create          = "PoolAssets.force_create"
	ExtName_PoolAssets_force_set_metadata    = "PoolAssets.force_set_metadata"
	ExtName_PoolAssets_force_transfer        = "PoolAssets.force_transfer"
	ExtName_PoolAssets_freeze                = "PoolAssets.freeze"
	ExtName_PoolAssets_freeze_asset          = "PoolAssets.freeze_asset"
	ExtName_PoolAssets_mint                  = "PoolAssets.mint"
	ExtName_PoolAssets_refund                = "PoolAssets.refund"
	ExtName_PoolAssets_refund_other          = "PoolAssets.refund_other"
	ExtName_PoolAssets_set_metadata          = "PoolAssets.set_metadata"
	ExtName_PoolAssets_set_min_balance       = "PoolAssets.set_min_balance"
	ExtName_PoolAssets_set_team              = "PoolAssets.set_team"
	ExtName_PoolAssets_start_destroy         = "PoolAssets.start_destroy"
	ExtName_PoolAssets_thaw                  = "PoolAssets.thaw"
	ExtName_PoolAssets_thaw_asset            = "PoolAssets.thaw_asset"
	ExtName_PoolAssets_touch                 = "PoolAssets.touch"
	ExtName_PoolAssets_touch_other           = "PoolAssets.touch_other"
	ExtName_PoolAssets_transfer              = "PoolAssets.transfer"
	ExtName_PoolAssets_transfer_all          = "PoolAssets.transfer_all"
	ExtName_PoolAssets_transfer_approved     = "PoolAssets.transfer_approved"
	ExtName_PoolAssets_transfer_keep_alive   = "PoolAssets.transfer_keep_alive"
	ExtName_PoolAssets_transfer_ownership    = "PoolAssets.transfer_ownership"

	// Preimage
	ExtName_Preimage_ensure_updated     = "Preimage.ensure_updated"
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

	// Reservoir
	ExtName_Reservoir_attend_evnet   = "Reservoir.attend_event"
	ExtName_Reservoir_create_event   = "Reservoir.create_event"
	ExtName_Reservoir_event_withdraw = "Reservoir.event_withdraw"
	ExtName_Reservoir_filling        = "Reservoir.filling"
	ExtName_Reservoir_store          = "Reservoir.store"
	ExtName_Reservoir_withdraw       = "Reservoir.withdraw"

	// Scheduler
	ExtName_Scheduler_cancel               = "Scheduler.cancel"
	ExtName_Scheduler_cancel_named         = "Scheduler.cancel_named"
	ExtName_Scheduler_cancel_retry         = "Scheduler.cancel_retry"
	ExtName_Scheduler_cancel_retry_named   = "Scheduler.cancel_retry_named"
	ExtName_Scheduler_schedule             = "Scheduler.schedule"
	ExtName_Scheduler_schedule_after       = "Scheduler.schedule_after"
	ExtName_Scheduler_schedule_named       = "Scheduler.schedule_named"
	ExtName_Scheduler_schedule_named_after = "Scheduler.schedule_named_after"
	ExtName_Scheduler_set_retry            = "Scheduler.set_retry"
	ExtName_Scheduler_set_retry_named      = "Scheduler.set_retry_named"

	// Session
	ExtName_Session_purge_keys = "Session.purge_keys"
	ExtName_Session_set_keys   = "Session.set_keys"

	// Sminer
	ExtName_Sminer_clear_miner_service        = "Sminer.clear_miner_service"
	ExtName_Sminer_decrease_declaration_space = "Sminer.decrease_declaration_space"
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
	ExtName_Sminer_set_facuet_whitelist       = "Sminer.set_facuet_whitelist"
	ExtName_Sminer_update_beneficiary         = "Sminer.update_beneficiary"
	ExtName_Sminer_update_endpoint            = "Sminer.update_endpoint"
	ExtName_Sminer_update_expender            = "Sminer.update_expender"

	// Staking
	ExtName_Staking_bond                       = "Staking.bond"
	ExtName_Staking_bond_extra                 = "Staking.bond_extra"
	ExtName_Staking_cancel_deferred_slash      = "Staking.cancel_deferred_slash"
	ExtName_Staking_chill                      = "Staking.chill"
	ExtName_Staking_chill_other                = "Staking.chill_other"
	ExtName_Staking_deprecate_controller_batch = "Staking.deprecate_controller_batch"
	ExtName_Staking_force_apply_min_commission = "Staking.force_apply_min_commission"
	ExtName_Staking_force_new_era              = "Staking.force_new_era"
	ExtName_Staking_force_new_era_always       = "Staking.force_new_era_always"
	ExtName_Staking_force_no_eras              = "Staking.force_no_eras"
	ExtName_Staking_force_unstake              = "Staking.force_unstake"
	ExtName_Staking_increase_validator_count   = "Staking.increase_validator_count"
	ExtName_Staking_kick                       = "Staking.kick"
	ExtName_Staking_nominate                   = "Staking.nominate"
	ExtName_Staking_payout_stakers             = "Staking.payout_stakers"
	ExtName_Staking_payout_stakers_by_page     = "Staking.payout_stakers_by_page"
	ExtName_Staking_reap_stash                 = "Staking.reap_stash"
	ExtName_Staking_rebond                     = "Staking.rebond"
	ExtName_Staking_restore_ledger             = "Staking.restore_ledger"
	ExtName_Staking_scale_validator_count      = "Staking.scale_validator_count"
	ExtName_Staking_set_controller             = "Staking.set_controller"
	ExtName_Staking_set_invulnerables          = "Staking.set_invulnerables"
	ExtName_Staking_set_min_commission         = "Staking.set_min_commission"
	ExtName_Staking_set_payee                  = "Staking.set_payee"
	ExtName_Staking_set_staking_configs        = "Staking.set_staking_configs"
	ExtName_Staking_set_validator_count        = "Staking.set_validator_count"
	ExtName_Staking_unbond                     = "Staking.unbond"
	ExtName_Staking_update_payee               = "Staking.update_payee"
	ExtName_Staking_validate                   = "Staking.validate"
	ExtName_Staking_withdraw_unbonded          = "Staking.withdraw_unbonded"

	// StateTrieMigration
	ExtName_StateTrieMigration_continue_migrate       = "StateTrieMigration.continue_migrate"
	ExtName_StateTrieMigration_control_auto_migration = "StateTrieMigration.control_auto_migration"
	ExtName_StateTrieMigration_force_set_progress     = "StateTrieMigration.force_set_progress"
	ExtName_StateTrieMigration_migrate_custom_child   = "StateTrieMigration.migrate_custom_child"
	ExtName_StateTrieMigration_migrate_custom_top     = "StateTrieMigration.migrate_custom_top"
	ExtName_StateTrieMigration_set_signed_max_limits  = "StateTrieMigration.set_signed_max_limits"

	// StorageHandler
	ExtName_StorageHandler_buy_consignment            = "StorageHandler.buy_consignment"
	ExtName_StorageHandler_cancel_consignment         = "StorageHandler.cancel_consignment"
	ExtName_StorageHandler_cancel_purchase_action     = "StorageHandler.cancel_purchase_action"
	ExtName_StorageHandler_clear_service_space        = "StorageHandler.clear_service_space"
	ExtName_StorageHandler_create_order               = "StorageHandler.create_order"
	ExtName_StorageHandler_define_update_price        = "StorageHandler.define_update_price"
	ExtName_StorageHandler_exec_consignment           = "StorageHandler.exec_consignment"
	ExtName_StorageHandler_exec_order                 = "StorageHandler.exec_order"
	ExtName_StorageHandler_expanding_territory        = "StorageHandler.expanding_territory"
	ExtName_StorageHandler_mint_territory             = "StorageHandler.mint_territory"
	ExtName_StorageHandler_reactivate_territory       = "StorageHandler.reactivate_territory"
	ExtName_StorageHandler_renewal_territory          = "StorageHandler.renewal_territory"
	ExtName_StorageHandler_territory_consignment      = "StorageHandler.territory_consignment"
	ExtName_StorageHandler_territory_grants           = "StorageHandler.territory_grants"
	ExtName_StorageHandler_territory_rename           = "StorageHandler.territory_rename"
	ExtName_StorageHandler_update_expired_exec        = "StorageHandler.update_expired_exec"
	ExtName_StorageHandler_update_price               = "StorageHandler.update_price"
	ExtName_StorageHandler_update_user_territory_life = "StorageHandler.update_user_territory_life"

	// Sudo
	ExtName_Sudo_remove_key            = "Sudo.remove_key"
	ExtName_Sudo_set_key               = "Sudo.set_key"
	ExtName_Sudo_sudo                  = "Sudo.sudo"
	ExtName_Sudo_sudo_as               = "Sudo.sudo_as"
	ExtName_Sudo_sudo_unchecked_weight = "Sudo.sudo_unchecked_weight"

	// System
	ExtName_System_apply_authorized_upgrade         = "System.apply_authorized_upgrade"
	ExtName_System_authorize_upgrade                = "System.authorize_upgrade"
	ExtName_System_authorize_upgrade_without_checks = "System.authorize_upgrade_without_checks"
	ExtName_System_kill_prefix                      = "System.kill_prefix"
	ExtName_System_kill_storage                     = "System.kill_storage"
	ExtName_System_remark                           = "System.remark"
	ExtName_System_remark_with_event                = "System.remark_with_event"
	ExtName_System_set_code                         = "System.set_code"
	ExtName_System_set_code_without_checks          = "System.set_code_without_checks"
	ExtName_System_set_heap_pages                   = "System.set_heap_pages"
	ExtName_System_set_storage                      = "System.set_storage"

	// TechnicalCommittee
	ExtName_TechnicalCommittee_close               = "TechnicalCommittee.close"
	ExtName_TechnicalCommittee_disapprove_proposal = "TechnicalCommittee.disapprove_proposal"
	ExtName_TechnicalCommittee_execute             = "TechnicalCommittee.execute"
	ExtName_TechnicalCommittee_propose             = "TechnicalCommittee.propose"
	ExtName_TechnicalCommittee_set_members         = "TechnicalCommittee.set_members"
	ExtName_TechnicalCommittee_vote                = "TechnicalCommittee.vote"

	// TeeWorker
	ExtName_TeeWorker_add_ceseal                 = "TeeWorker.add_ceseal"
	ExtName_TeeWorker_apply_master_key           = "TeeWorker.apply_master_key"
	ExtName_TeeWorker_change_first_holder        = "TeeWorker.change_first_holder"
	ExtName_TeeWorker_clear_master_key           = "TeeWorker.clear_master_key"
	ExtName_TeeWorker_force_clear_tee            = "TeeWorker.force_clear_tee"
	ExtName_TeeWorker_force_register_worker      = "TeeWorker.force_register_worker"
	ExtName_TeeWorker_launch_master_key          = "TeeWorker.launch_master_key"
	ExtName_TeeWorker_migration_last_work        = "TeeWorker.migration_last_work"
	ExtName_TeeWorker_patch_clear_invalid_tee    = "TeeWorker.patch_clear_invalid_tee"
	ExtName_TeeWorker_patch_clear_not_work_tee   = "TeeWorker.patch_clear_not_work_tee"
	ExtName_TeeWorker_refresh_tee_status         = "TeeWorker.refresh_tee_status"
	ExtName_TeeWorker_register_worker            = "TeeWorker.register_worker"
	ExtName_TeeWorker_register_worker_v2         = "TeeWorker.register_worker_v2"
	ExtName_TeeWorker_remove_ceseal              = "TeeWorker.remove_ceseal"
	ExtName_TeeWorker_set_minimum_ceseal_version = "TeeWorker.set_minimum_ceseal_version"
	ExtName_TeeWorker_set_note_stalled           = "TeeWorker.set_note_stalled"
	ExtName_TeeWorker_update_worker_endpoint     = "TeeWorker.update_worker_endpoint"

	// Timestamp
	ExtName_Timestamp_set = "Timestamp.set"

	// TransactionStorage
	ExtName_TransactionStorage_check_proof = "TransactionStorage.check_proof"
	ExtName_TransactionStorage_renew       = "TransactionStorage.renew"
	ExtName_TransactionStorage_store       = "TransactionStorage.store"

	// Treasury
	ExtName_Treasury_check_status    = "Treasury.check_status"
	ExtName_Treasury_payout          = "Treasury.payout"
	ExtName_Treasury_remove_approval = "Treasury.remove_approval"
	ExtName_Treasury_spend           = "Treasury.spend"
	ExtName_Treasury_spend_local     = "Treasury.spend_local"
	ExtName_Treasury_void_spend      = "Treasury.void_spend"

	// Utility
	ExtName_Utility_as_derivative = "Utility.as_derivative"
	ExtName_Utility_batch         = "Utility.batch"
	ExtName_Utility_batch_all     = "Utility.batch_all"
	ExtName_Utility_dispatch_as   = "Utility.dispatch_as"
	ExtName_Utility_force_batch   = "Utility.force_batch"
	ExtName_Utility_with_weight   = "Utility.with_weight"

	// VoterList
	ExtName_VoterList_put_in_front_of       = "VoterList.put_in_front_of"
	ExtName_VoterList_put_in_front_of_other = "VoterList.put_in_front_of_other"
	ExtName_VoterList_rebag                 = "VoterList.rebag"
)

// InitExtrinsicsName initialises all transaction names
//
// Return:
//   - error: error message
//
// Note:
//   - If you need to parse all transaction events, you need to call this function.
func (c *ChainClient) InitExtrinsicsName() error {
	ExtrinsicsName = make(map[types.CallIndex]string, 0)

	// AssetConversion
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_AssetConversion_add_liquidity); err == nil {
		ExtrinsicsName[callIndex] = ExtName_AssetConversion_add_liquidity
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_AssetConversion_create_pool); err == nil {
		ExtrinsicsName[callIndex] = ExtName_AssetConversion_create_pool
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_AssetConversion_remove_liquidity); err == nil {
		ExtrinsicsName[callIndex] = ExtName_AssetConversion_remove_liquidity
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_AssetConversion_swap_exact_tokens_for_tokens); err == nil {
		ExtrinsicsName[callIndex] = ExtName_AssetConversion_swap_exact_tokens_for_tokens
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_AssetConversion_swap_tokens_for_exact_tokens); err == nil {
		ExtrinsicsName[callIndex] = ExtName_AssetConversion_swap_tokens_for_exact_tokens
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_AssetConversion_stouch); err == nil {
		ExtrinsicsName[callIndex] = ExtName_AssetConversion_stouch
	} else {
		return err
	}

	// AssetRate
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_AssetRate_create); err == nil {
		ExtrinsicsName[callIndex] = ExtName_AssetRate_create
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_AssetRate_remove); err == nil {
		ExtrinsicsName[callIndex] = ExtName_AssetRate_remove
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_AssetRate_update); err == nil {
		ExtrinsicsName[callIndex] = ExtName_AssetRate_update
	} else {
		return err
	}

	// Assets
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_approve_transfer); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_approve_transfer
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_block); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_block
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_burn); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_burn
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_cancel_approval); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_cancel_approval
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_clear_metadata); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_clear_metadata
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_create); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_create
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_destroy_accounts); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_destroy_accounts
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_destroy_approvals); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_destroy_approvals
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_finish_destroy); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_finish_destroy
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_force_asset_status); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_force_asset_status
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_force_cancel_approval); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_force_cancel_approval
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_force_clear_metadata); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_force_clear_metadata
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_force_create); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_force_create
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_force_set_metadata); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_force_set_metadata
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_force_transfer); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_force_transfer
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_freeze); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_freeze
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_freeze_asset); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_freeze_asset
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_mint); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_mint
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_refund); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_refund
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_refund_other); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_refund_other
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_set_metadata); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_set_metadata
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_set_min_balance); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_set_min_balance
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_set_team); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_set_team
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_start_destroy); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_start_destroy
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_thaw); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_thaw
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_thaw_asset); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_thaw_asset
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_touch); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_touch
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_touch_other); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_touch_other
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_transfer); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_transfer
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_transfer_all); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_transfer_all
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_transfer_approved); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_transfer_approved
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_transfer_keep_alive); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_transfer_keep_alive
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Assets_transfer_ownership); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Assets_transfer_ownership
	} else {
		return err
	}

	// Audit
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Audit_point_miner_challenge); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Audit_point_miner_challenge
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Audit_submit_idle_proof); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Audit_submit_idle_proof
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Audit_submit_service_proof); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Audit_submit_service_proof
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Audit_submit_verify_idle_result); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Audit_submit_verify_idle_result
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Audit_submit_verify_service_result); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Audit_submit_verify_service_result
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Audit_test_update_clear_slip); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Audit_test_update_clear_slip
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Audit_test_update_verify_slip); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Audit_test_update_verify_slip
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Audit_update_counted_clear); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Audit_update_counted_clear
	} else {
		return err
	}

	// Babe
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Babe_plan_config_change); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Babe_plan_config_change
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Babe_report_equivocation); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Babe_report_equivocation
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Babe_report_equivocation_unsigned); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Babe_report_equivocation_unsigned
	} else {
		return err
	}

	// Balances
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Balances_burn); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Balances_burn
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Balances_force_adjust_total_issuance); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Balances_force_adjust_total_issuance
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Balances_force_set_balance); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Balances_force_set_balance
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Balances_force_transfer); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Balances_force_transfer
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Balances_force_unreserve); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Balances_force_unreserve
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Balances_transfer_all); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Balances_transfer_all
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Balances_transfer_allow_death); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Balances_transfer_allow_death
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Balances_transferKeepAlive); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Balances_transferKeepAlive
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Balances_upgrade_accounts); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Balances_upgrade_accounts
	} else {
		return err
	}

	// BaseFee
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_BaseFee_set_base_fee_per_gas); err == nil {
		ExtrinsicsName[callIndex] = ExtName_BaseFee_set_base_fee_per_gas
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_BaseFee_set_elasticity); err == nil {
		ExtrinsicsName[callIndex] = ExtName_BaseFee_set_elasticity
	} else {
		return err
	}

	// Cacher
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Cacher_logout); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Cacher_logout
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Cacher_pay); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Cacher_pay
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Cacher_register); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Cacher_register
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Cacher_update); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Cacher_update
	} else {
		return err
	}

	// CesMq
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_CesMq_force_push_pallet_message); err == nil {
		ExtrinsicsName[callIndex] = ExtName_CesMq_force_push_pallet_message
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_CesMq_push_message); err == nil {
		ExtrinsicsName[callIndex] = ExtName_CesMq_push_message
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_CesMq_sync_offchain_message); err == nil {
		ExtrinsicsName[callIndex] = ExtName_CesMq_sync_offchain_message
	} else {
		return err
	}

	// CessTreasury
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_CessTreasury_pid_burn_funds); err == nil {
		ExtrinsicsName[callIndex] = ExtName_CessTreasury_pid_burn_funds
	} else {
		return err
	}

	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_CessTreasury_pid_send_funds); err == nil {
		ExtrinsicsName[callIndex] = ExtName_CessTreasury_pid_send_funds
	} else {
		return err
	}

	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_CessTreasury_send_funds_to_pid); err == nil {
		ExtrinsicsName[callIndex] = ExtName_CessTreasury_send_funds_to_pid
	} else {
		return err
	}

	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_CessTreasury_send_funds_to_sid); err == nil {
		ExtrinsicsName[callIndex] = ExtName_CessTreasury_send_funds_to_sid
	} else {
		return err
	}

	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_CessTreasury_sid_burn_funds); err == nil {
		ExtrinsicsName[callIndex] = ExtName_CessTreasury_sid_burn_funds
	} else {
		return err
	}

	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_CessTreasury_sid_send_funds); err == nil {
		ExtrinsicsName[callIndex] = ExtName_CessTreasury_sid_send_funds
	} else {
		return err
	}

	// Contracts
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Contracts_call); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Contracts_call
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Contracts_call_old_weight); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Contracts_call_old_weight
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Contracts_instantiate); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Contracts_instantiate
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Contracts_instantiate_old_weight); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Contracts_instantiate_old_weight
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Contracts_instantiate_with_code); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Contracts_instantiate_with_code
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Contracts_instantiate_with_code_old_weight); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Contracts_instantiate_with_code_old_weight
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Contracts_migrate); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Contracts_migrate
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Contracts_remove_code); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Contracts_remove_code
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Contracts_set_code); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Contracts_set_code
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Contracts_upload_code); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Contracts_upload_code
	} else {
		return err
	}

	// Council
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Council_close); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Council_close
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Council_disapprove_proposal); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Council_disapprove_proposal
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Council_execute); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Council_execute
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Council_propose); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Council_propose
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Council_set_members); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Council_set_members
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Council_vote); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Council_vote
	} else {
		return err
	}

	// ElectionProviderMultiPhase
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_ElectionProviderMultiPhase_governance_fallback); err == nil {
		ExtrinsicsName[callIndex] = ExtName_ElectionProviderMultiPhase_governance_fallback
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_ElectionProviderMultiPhase_set_emergency_election_result); err == nil {
		ExtrinsicsName[callIndex] = ExtName_ElectionProviderMultiPhase_set_emergency_election_result
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_ElectionProviderMultiPhase_set_minimum_untrusted_score); err == nil {
		ExtrinsicsName[callIndex] = ExtName_ElectionProviderMultiPhase_set_minimum_untrusted_score
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_ElectionProviderMultiPhase_submit); err == nil {
		ExtrinsicsName[callIndex] = ExtName_ElectionProviderMultiPhase_submit
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_ElectionProviderMultiPhase_submit_unsigned); err == nil {
		ExtrinsicsName[callIndex] = ExtName_ElectionProviderMultiPhase_submit_unsigned
	} else {
		return err
	}

	// Ethereum
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Ethereum_transact); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Ethereum_transact
	} else {
		return err
	}

	// Evm
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Evm_call); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Evm_call
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Evm_create); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Evm_create
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Evm_create2); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Evm_create2
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Evm_withdraw); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Evm_withdraw
	} else {
		return err
	}

	// EvmAccountMapping
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_EvmAccountMapping_meta_call); err == nil {
		ExtrinsicsName[callIndex] = ExtName_EvmAccountMapping_meta_call
	} else {
		return err
	}

	// FastUnstake
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FastUnstake_control); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FastUnstake_control
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FastUnstake_deregister); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FastUnstake_deregister
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FastUnstake_register_fast_unstake); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FastUnstake_register_fast_unstake
	} else {
		return err
	}

	// FileBank
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_calculate_report); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_calculate_report
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_cert_idle_space); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_cert_idle_space
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_claim_restoral_noexist_order); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_claim_restoral_noexist_order
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_claim_restoral_order); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_claim_restoral_order
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_delete_file); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_delete_file
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_generate_restoral_order); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_generate_restoral_order
	} else {
		return err
	}

	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_replace_idle_space); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_replace_idle_space
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_restoral_order_complete); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_restoral_order_complete
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_root_clear_file); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_root_clear_file
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_transfer_report); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_transfer_report
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_upload_declaration); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_upload_declaration
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_territory_file_delivery); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_territory_file_delivery
	} else {
		return err
	}

	// Grandpa
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Grandpa_note_stalled); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Grandpa_note_stalled
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Grandpa_report_equivocation); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Grandpa_report_equivocation
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Grandpa_report_equivocation_unsigned); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Grandpa_report_equivocation_unsigned
	} else {
		return err
	}

	// ImOnline
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_ImOnline_heartbeat); err == nil {
		ExtrinsicsName[callIndex] = ExtName_ImOnline_heartbeat
	} else {
		return err
	}

	// Indices
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Indices_claim); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Indices_claim
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Indices_force_transfer); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Indices_force_transfer
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Indices_free); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Indices_free
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Indices_freeze); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Indices_freeze
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Indices_transfer); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Indices_transfer
	} else {
		return err
	}

	// MultiBlockMigrations
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_MultiBlockMigrations_clear_historic); err == nil {
		ExtrinsicsName[callIndex] = ExtName_MultiBlockMigrations_clear_historic
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_MultiBlockMigrations_force_onboard_mbms); err == nil {
		ExtrinsicsName[callIndex] = ExtName_MultiBlockMigrations_force_onboard_mbms
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_MultiBlockMigrations_force_set_active_cursor); err == nil {
		ExtrinsicsName[callIndex] = ExtName_MultiBlockMigrations_force_set_active_cursor
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_MultiBlockMigrations_force_set_cursor); err == nil {
		ExtrinsicsName[callIndex] = ExtName_MultiBlockMigrations_force_set_cursor
	} else {
		return err
	}

	// Multisig
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Multisig_approve_as_multi); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Multisig_approve_as_multi
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Multisig_as_multi); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Multisig_as_multi
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Multisig_as_multi_threshold1); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Multisig_as_multi_threshold1
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Multisig_cancel_as_multi); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Multisig_cancel_as_multi
	} else {
		return err
	}

	// Oss
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Oss_authorize); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Oss_authorize
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Oss_cancel_authorize); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Oss_cancel_authorize
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Oss_destroy); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Oss_destroy
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Oss_evm_proxy_authorzie); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Oss_evm_proxy_authorzie
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Oss_proxy_authorzie); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Oss_proxy_authorzie
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Oss_register); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Oss_register
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Oss_update); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Oss_update
	} else {
		return err
	}

	// Parameters
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Parameters_set_parameter); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Parameters_set_parameter
	} else {
		return err
	}

	// PoolAssets
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_approve_transfer); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_approve_transfer
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_block); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_block
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_burn); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_burn
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_cancel_approval); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_cancel_approval
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_clear_metadata); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_clear_metadata
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_create); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_create
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_destroy_accounts); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_destroy_accounts
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_destroy_approvals); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_destroy_approvals
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_finish_destroy); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_finish_destroy
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_force_asset_status); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_force_asset_status
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_force_cancel_approval); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_force_cancel_approval
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_force_clear_metadata); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_force_clear_metadata
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_force_create); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_force_create
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_force_set_metadata); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_force_set_metadata
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_force_transfer); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_force_transfer
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_freeze); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_freeze
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_freeze_asset); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_freeze_asset
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_mint); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_mint
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_refund); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_refund
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_refund_other); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_refund_other
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_set_metadata); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_set_metadata
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_set_min_balance); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_set_min_balance
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_set_team); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_set_team
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_start_destroy); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_start_destroy
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_thaw); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_thaw
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_thaw_asset); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_thaw_asset
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_touch); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_touch
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_touch_other); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_touch_other
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_transfer); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_transfer
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_transfer_all); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_transfer_all
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_transfer_approved); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_transfer_approved
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_transfer_keep_alive); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_transfer_keep_alive
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_PoolAssets_transfer_ownership); err == nil {
		ExtrinsicsName[callIndex] = ExtName_PoolAssets_transfer_ownership
	} else {
		return err
	}

	// Preimage
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Preimage_ensure_updated); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Preimage_ensure_updated
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Preimage_note_preimage); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Preimage_note_preimage
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Preimage_request_preimage); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Preimage_request_preimage
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Preimage_unnote_preimage); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Preimage_unnote_preimage
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Preimage_unrequest_preimage); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Preimage_unrequest_preimage
	} else {
		return err
	}

	// Proxy
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Proxy_add_proxy); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Proxy_add_proxy
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Proxy_announce); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Proxy_announce
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Proxy_create_pure); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Proxy_create_pure
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Proxy_kill_pure); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Proxy_kill_pure
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Proxy_proxy); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Proxy_proxy
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Proxy_proxy_announced); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Proxy_proxy_announced
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Proxy_reject_announcement); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Proxy_reject_announcement
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Proxy_remove_announcement); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Proxy_remove_announcement
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Proxy_remove_proxies); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Proxy_remove_proxies
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Proxy_remove_proxy); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Proxy_remove_proxy
	} else {
		return err
	}

	// Reservoir
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Reservoir_attend_evnet); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Reservoir_attend_evnet
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Reservoir_create_event); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Reservoir_create_event
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Reservoir_event_withdraw); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Reservoir_event_withdraw
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Reservoir_filling); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Reservoir_filling
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Reservoir_store); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Reservoir_store
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Reservoir_withdraw); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Reservoir_withdraw
	} else {
		return err
	}

	// Scheduler
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Scheduler_cancel); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Scheduler_cancel
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Scheduler_cancel_named); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Scheduler_cancel_named
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Scheduler_cancel_retry); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Scheduler_cancel_retry
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Scheduler_cancel_retry_named); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Scheduler_cancel_retry_named
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Scheduler_schedule); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Scheduler_schedule
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Scheduler_schedule_after); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Scheduler_schedule_after
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Scheduler_schedule_named); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Scheduler_schedule_named
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Scheduler_schedule_named_after); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Scheduler_schedule_named_after
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Scheduler_set_retry); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Scheduler_set_retry
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Scheduler_set_retry_named); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Scheduler_set_retry_named
	} else {
		return err
	}

	// Session
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Session_purge_keys); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Session_purge_keys
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Session_set_keys); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Session_set_keys
	} else {
		return err
	}

	// Sminer
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_clear_miner_service); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_clear_miner_service
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_decrease_declaration_space); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_decrease_declaration_space
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_faucet); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_faucet
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_faucet_top_up); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_faucet_top_up
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_increase_collateral); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_increase_collateral
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_increase_declaration_space); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_increase_declaration_space
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_miner_exit); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_miner_exit
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_miner_exit_prep); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_miner_exit_prep
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_miner_withdraw); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_miner_withdraw
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_receive_reward); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_receive_reward
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_register_pois_key); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_register_pois_key
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_regnstk); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_regnstk
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_regnstk_assign_staking); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_regnstk_assign_staking
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_set_facuet_whitelist); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_set_facuet_whitelist
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_update_beneficiary); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_update_beneficiary
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_update_endpoint); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_update_endpoint
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_update_expender); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_update_expender
	} else {
		return err
	}

	// Staking
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_bond); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_bond
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_bond_extra); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_bond_extra
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_cancel_deferred_slash); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_cancel_deferred_slash
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_chill); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_chill
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_chill_other); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_chill_other
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_deprecate_controller_batch); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_deprecate_controller_batch
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_force_apply_min_commission); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_force_apply_min_commission
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_force_new_era); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_force_new_era
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_force_new_era_always); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_force_new_era_always
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_force_no_eras); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_force_no_eras
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_force_unstake); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_force_unstake
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_increase_validator_count); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_increase_validator_count
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_kick); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_kick
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_nominate); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_nominate
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_payout_stakers); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_payout_stakers
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_payout_stakers_by_page); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_payout_stakers_by_page
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_reap_stash); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_reap_stash
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_rebond); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_rebond
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_restore_ledger); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_restore_ledger
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_scale_validator_count); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_scale_validator_count
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_set_controller); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_set_controller
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_set_invulnerables); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_set_invulnerables
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_set_min_commission); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_set_min_commission
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_set_payee); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_set_payee
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_set_staking_configs); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_set_staking_configs
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_set_validator_count); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_set_validator_count
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_unbond); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_unbond
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_update_payee); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_update_payee
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_validate); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_validate
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Staking_withdraw_unbonded); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Staking_withdraw_unbonded
	} else {
		return err
	}

	// StateTrieMigration
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StateTrieMigration_continue_migrate); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StateTrieMigration_continue_migrate
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StateTrieMigration_control_auto_migration); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StateTrieMigration_control_auto_migration
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StateTrieMigration_force_set_progress); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StateTrieMigration_force_set_progress
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StateTrieMigration_migrate_custom_child); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StateTrieMigration_migrate_custom_child
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StateTrieMigration_migrate_custom_top); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StateTrieMigration_migrate_custom_top
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StateTrieMigration_set_signed_max_limits); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StateTrieMigration_set_signed_max_limits
	} else {
		return err
	}

	// StorageHandler
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StorageHandler_buy_consignment); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StorageHandler_buy_consignment
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StorageHandler_cancel_consignment); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StorageHandler_cancel_consignment
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StorageHandler_cancel_purchase_action); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StorageHandler_cancel_purchase_action
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StorageHandler_clear_service_space); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StorageHandler_clear_service_space
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StorageHandler_create_order); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StorageHandler_create_order
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StorageHandler_define_update_price); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StorageHandler_define_update_price
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StorageHandler_exec_consignment); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StorageHandler_exec_consignment
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StorageHandler_exec_order); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StorageHandler_exec_order
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StorageHandler_expanding_territory); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StorageHandler_expanding_territory
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StorageHandler_mint_territory); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StorageHandler_mint_territory
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StorageHandler_reactivate_territory); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StorageHandler_reactivate_territory
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StorageHandler_renewal_territory); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StorageHandler_renewal_territory
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StorageHandler_territory_consignment); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StorageHandler_territory_consignment
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StorageHandler_territory_grants); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StorageHandler_territory_grants
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StorageHandler_territory_rename); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StorageHandler_territory_rename
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StorageHandler_update_expired_exec); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StorageHandler_update_expired_exec
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StorageHandler_update_price); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StorageHandler_update_price
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_StorageHandler_update_user_territory_life); err == nil {
		ExtrinsicsName[callIndex] = ExtName_StorageHandler_update_user_territory_life
	} else {
		return err
	}

	// Sudo
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sudo_remove_key); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sudo_remove_key
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sudo_set_key); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sudo_set_key
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sudo_sudo); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sudo_sudo
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sudo_sudo_as); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sudo_sudo_as
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sudo_sudo_unchecked_weight); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sudo_sudo_unchecked_weight
	} else {
		return err
	}

	// System
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_System_apply_authorized_upgrade); err == nil {
		ExtrinsicsName[callIndex] = ExtName_System_apply_authorized_upgrade
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_System_authorize_upgrade); err == nil {
		ExtrinsicsName[callIndex] = ExtName_System_authorize_upgrade
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_System_authorize_upgrade_without_checks); err == nil {
		ExtrinsicsName[callIndex] = ExtName_System_authorize_upgrade_without_checks
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_System_kill_prefix); err == nil {
		ExtrinsicsName[callIndex] = ExtName_System_kill_prefix
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_System_kill_storage); err == nil {
		ExtrinsicsName[callIndex] = ExtName_System_kill_storage
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_System_remark); err == nil {
		ExtrinsicsName[callIndex] = ExtName_System_remark
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_System_remark_with_event); err == nil {
		ExtrinsicsName[callIndex] = ExtName_System_remark_with_event
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_System_set_code); err == nil {
		ExtrinsicsName[callIndex] = ExtName_System_set_code
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_System_set_code_without_checks); err == nil {
		ExtrinsicsName[callIndex] = ExtName_System_set_code_without_checks
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_System_set_heap_pages); err == nil {
		ExtrinsicsName[callIndex] = ExtName_System_set_heap_pages
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_System_set_storage); err == nil {
		ExtrinsicsName[callIndex] = ExtName_System_set_storage
	} else {
		return err
	}

	// TechnicalCommittee
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TechnicalCommittee_close); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TechnicalCommittee_close
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TechnicalCommittee_disapprove_proposal); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TechnicalCommittee_disapprove_proposal
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TechnicalCommittee_execute); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TechnicalCommittee_execute
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TechnicalCommittee_propose); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TechnicalCommittee_propose
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TechnicalCommittee_set_members); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TechnicalCommittee_set_members
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TechnicalCommittee_vote); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TechnicalCommittee_vote
	} else {
		return err
	}

	// TeeWorker
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TeeWorker_add_ceseal); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TeeWorker_add_ceseal
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TeeWorker_apply_master_key); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TeeWorker_apply_master_key
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TeeWorker_change_first_holder); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TeeWorker_change_first_holder
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TeeWorker_clear_master_key); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TeeWorker_clear_master_key
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TeeWorker_force_clear_tee); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TeeWorker_force_clear_tee
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TeeWorker_force_register_worker); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TeeWorker_force_register_worker
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TeeWorker_launch_master_key); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TeeWorker_launch_master_key
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TeeWorker_migration_last_work); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TeeWorker_migration_last_work
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TeeWorker_patch_clear_invalid_tee); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TeeWorker_patch_clear_invalid_tee
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TeeWorker_patch_clear_not_work_tee); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TeeWorker_patch_clear_not_work_tee
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TeeWorker_refresh_tee_status); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TeeWorker_refresh_tee_status
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TeeWorker_register_worker); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TeeWorker_register_worker
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TeeWorker_register_worker_v2); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TeeWorker_register_worker_v2
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TeeWorker_remove_ceseal); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TeeWorker_remove_ceseal
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TeeWorker_set_minimum_ceseal_version); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TeeWorker_set_minimum_ceseal_version
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TeeWorker_set_note_stalled); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TeeWorker_set_note_stalled
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TeeWorker_update_worker_endpoint); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TeeWorker_update_worker_endpoint
	} else {
		return err
	}

	// Timestamp
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Timestamp_set); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Timestamp_set
	} else {
		return err
	}

	// TransactionStorage
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TransactionStorage_check_proof); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TransactionStorage_check_proof
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TransactionStorage_renew); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TransactionStorage_renew
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_TransactionStorage_store); err == nil {
		ExtrinsicsName[callIndex] = ExtName_TransactionStorage_store
	} else {
		return err
	}

	// Treasury
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Treasury_check_status); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Treasury_check_status
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Treasury_payout); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Treasury_payout
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Treasury_remove_approval); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Treasury_remove_approval
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Treasury_spend); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Treasury_spend
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Treasury_spend_local); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Treasury_spend_local
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Treasury_void_spend); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Treasury_void_spend
	} else {
		return err
	}

	// Utility
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Utility_as_derivative); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Utility_as_derivative
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Utility_batch); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Utility_batch
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Utility_batch_all); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Utility_batch_all
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Utility_dispatch_as); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Utility_dispatch_as
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Utility_force_batch); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Utility_force_batch
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Utility_with_weight); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Utility_with_weight
	} else {
		return err
	}

	// VoterList
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_VoterList_put_in_front_of); err == nil {
		ExtrinsicsName[callIndex] = ExtName_VoterList_put_in_front_of
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_VoterList_put_in_front_of_other); err == nil {
		ExtrinsicsName[callIndex] = ExtName_VoterList_put_in_front_of_other
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_VoterList_rebag); err == nil {
		ExtrinsicsName[callIndex] = ExtName_VoterList_rebag
	} else {
		return err
	}
	return nil
}

// InitExtrinsicsNameForMiner initialises all transaction required by the storage miner
//
// Return:
//   - error: error message
//
// Note:
//   - The storage miner program needs to call this function, otherwise the transaction event cannot be parsed.
func (c *ChainClient) InitExtrinsicsNameForMiner() error {
	ExtrinsicsName = make(map[types.CallIndex]string, 0)

	// Audit
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Audit_submit_idle_proof); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Audit_submit_idle_proof
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Audit_submit_service_proof); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Audit_submit_service_proof
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Audit_submit_verify_idle_result); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Audit_submit_verify_idle_result
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Audit_submit_verify_service_result); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Audit_submit_verify_service_result
	} else {
		return err
	}

	// FileBank
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_calculate_report); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_calculate_report
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_cert_idle_space); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_cert_idle_space
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_claim_restoral_noexist_order); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_claim_restoral_noexist_order
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_claim_restoral_order); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_claim_restoral_order
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_generate_restoral_order); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_generate_restoral_order
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_replace_idle_space); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_replace_idle_space
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_restoral_order_complete); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_restoral_order_complete
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_transfer_report); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_transfer_report
	} else {
		return err
	}

	// Sminer
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_decrease_declaration_space); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_decrease_declaration_space
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_faucet); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_faucet
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_faucet_top_up); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_faucet_top_up
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_increase_collateral); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_increase_collateral
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_increase_declaration_space); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_increase_declaration_space
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_miner_exit); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_miner_exit
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_miner_exit_prep); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_miner_exit_prep
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_miner_withdraw); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_miner_withdraw
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_receive_reward); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_receive_reward
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_register_pois_key); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_register_pois_key
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_regnstk); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_regnstk
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_regnstk_assign_staking); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_regnstk_assign_staking
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_update_beneficiary); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_update_beneficiary
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Sminer_update_endpoint); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Sminer_update_endpoint
	} else {
		return err
	}

	// Timestamp
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Timestamp_set); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Timestamp_set
	} else {
		return err
	}

	return nil
}

// InitExtrinsicsNameForOSS initialises all transaction required by the deoss
//
// Return:
//   - error: error message
//
// Note:
//   - The deoss program needs to call this function, otherwise the transaction event cannot be parsed.
func (c *ChainClient) InitExtrinsicsNameForOSS() error {
	ExtrinsicsName = make(map[types.CallIndex]string, 0)

	// FileBank
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_delete_file); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_delete_file
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_upload_declaration); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_upload_declaration
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_FileBank_territory_file_delivery); err == nil {
		ExtrinsicsName[callIndex] = ExtName_FileBank_territory_file_delivery
	} else {
		return err
	}

	// Oss
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Oss_destroy); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Oss_destroy
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Oss_register); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Oss_register
	} else {
		return err
	}
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Oss_update); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Oss_update
	} else {
		return err
	}

	// Timestamp
	if callIndex, err := c.GetMetadata().FindCallIndex(ExtName_Timestamp_set); err == nil {
		ExtrinsicsName[callIndex] = ExtName_Timestamp_set
	} else {
		return err
	}
	return nil
}
