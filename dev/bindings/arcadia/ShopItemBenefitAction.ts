// This file was generated by [ts-rs](https://github.com/Aleph-Alpha/ts-rs). Do not edit this file manually.

export type ShopItemBenefitAction = "List" | { "Create": { 
/**
 * The ID of the shop item benefit
 */
id: string, 
/**
 * The friendly name of the shop item benefit
 */
name: string, 
/**
 * The description of the shop item benefit
 */
description: string, 
/**
 * The target types the benefit can be applied to
 */
target_types: Array<string>, } } | { "Edit": { 
/**
 * The ID of the shop item benefit
 */
id: string, 
/**
 * The friendly name of the shop item benefit
 */
name: string, 
/**
 * The description of the shop item benefit
 */
description: string, 
/**
 * The target types the benefit can be applied to
 */
target_types: Array<string>, } } | { "Delete": { 
/**
 * The ID of the shop item benefit
 */
id: string, } };