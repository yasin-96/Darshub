import type { Auth0UserExtendedObject } from "../user/types";
import type { Auth0UserStats } from "./types";

export interface Auth0UserManagementDetails {
  users: Array<Auth0UserExtendedObject>,
  stats: Auth0UserStats
}