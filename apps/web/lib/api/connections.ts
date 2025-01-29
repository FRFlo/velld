import { Connection, ConnectionForm } from "@/types/connection";
import { apiRequest } from "@/lib/api-client";

export async function getConnections(): Promise<Connection[]> {
  return apiRequest<Connection[]>("/api/connections");
}

export async function testConnection(connection: ConnectionForm) {
  return apiRequest("/api/connections/test", {
    method: "POST",
    body: JSON.stringify(connection),
  });
}

export async function saveConnection(connection: ConnectionForm) {
  return apiRequest<Connection>("/api/connections", {
    method: "POST",
    body: JSON.stringify(connection),
  });
}