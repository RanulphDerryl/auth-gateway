// types.ts
import { Reflector } from '@nestjs/core';

// AuthGatewayConfig interface
export interface AuthGatewayConfig {
  jwt: {
    secretKey: string;
    expiresIn: string;
  };
  database: {
    url: string;
    username: string;
    password: string;
    options?: any;
  };
  security: {
    allowedOrigins: string[];
  };
}

// AuthUser interface
export interface AuthUser {
  id: string;
  username: string;
  email: string;
  roles: string[];
}

// AuthUserRequest interface
export interface AuthUserRequest {
  username: string;
  password: string;
}

// AuthUserResponse interface
export interface AuthUserResponse {
  id: string;
  username: string;
  email: string;
  roles: string[];