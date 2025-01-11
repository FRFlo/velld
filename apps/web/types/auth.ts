import { Base } from "./base";
export interface Auth {
  username: string;
  password: string;
}

export interface AuthResponse extends Base<{ token: string }> {
  token: string;
}

export interface Profile extends Base<{ username: string }> {
  username: string;
}