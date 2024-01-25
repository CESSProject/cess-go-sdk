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
	ExtName_CessTreasury_pid_burn_funds = "CessTreasury.pid_burn_funds"
	ExtName_CessTreasury_pid_send_funds = "CessTreasury.pid_send_funds"
	// ExtName_CessTreasury_pid_burn_funds = "CessTreasury.send_funds_to_tid"
	// ExtName_CessTreasury_pid_burn_funds = "CessTreasury.pid_burn_funds"
	// ExtName_CessTreasury_pid_burn_funds = "CessTreasury.pid_burn_funds"
	// ExtName_CessTreasury_pid_burn_funds = "CessTreasury.pid_burn_funds"

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
