// This file was generated by [ts-rs](https://github.com/Aleph-Alpha/ts-rs). Do not edit this file manually.
import type { PlatformUser } from "./PlatformUser";
import type { StaffDisciplinary } from "./StaffDisciplinary";
import type { StaffPosition } from "./StaffPosition";

export type StaffMember = { 
/**
 * The id of the user
 */
user_id: string, 
/**
 * The user object of the staff member
 */
user: PlatformUser, 
/**
 * The positions of the staff member
 */
positions: Array<StaffPosition>, 
/**
 * The disciplinary actions recieved by the member
 */
disciplinaries: Array<StaffDisciplinary>, 
/**
 * The permission overrides of the staff member
 */
perm_overrides: Array<string>, 
/**
 * The resolved permissions available to the member
 */
resolved_perms: Array<string>, 
/**
 * Whether or not the member is 'frozen' and cannot be updated in resyncs
 */
no_autosync: boolean, 
/**
 * Whether or not the member is 'known' to be 'unaccounted' for
 */
unaccounted: boolean, 
/**
 * Whether or not the members MFA is verified or not
 */
mfa_verified: boolean, 
/**
 * When the staff member was created/added
 */
created_at: string, };