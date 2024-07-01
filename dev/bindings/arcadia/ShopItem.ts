// This file was generated by [ts-rs](https://github.com/Aleph-Alpha/ts-rs). Do not edit this file manually.

/**
 * Shop items are items that can be purchased by users on the shop
 */
export type ShopItem = { 
/**
 * The ID of the shop item
 */
id: string, 
/**
 * The friendly name of the shop item
 */
name: string, 
/**
 * The description of the shop item
 */
description: string, 
/**
 * The cents the shop item costs
 */
cents: number, 
/**
 * The target type
 */
target_types: Array<string>, 
/**
 * The benefits of the shop item
 */
benefits: Array<string>, 
/**
 * The number of hours the shop item lasts for
 */
duration: number, 
/**
 * The time the shop item was created
 */
created_at: string, 
/**
 * The time the shop item was last updated
 */
last_updated: string, 
/**
 * Who created the shop item
 */
created_by: string, 
/**
 * Who last updated the shop item
 */
updated_by: string, };