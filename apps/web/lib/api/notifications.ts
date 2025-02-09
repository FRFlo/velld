import { NotificationResponse } from "@/types/notification";
import { apiRequest } from "@/lib/api-client";

export async function getNotifications(): Promise<NotificationResponse> {
  return apiRequest<NotificationResponse>("/api/notifications");
}

export async function markAsRead(notificationId: string) {
  return apiRequest("/api/notifications/mark-read", {
    method: "POST",
    body: JSON.stringify({ notification_id: notificationId }),
  });
}
