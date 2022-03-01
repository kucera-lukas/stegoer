import { UnstyledButton } from "@mantine/core";
import { useNotifications } from "@mantine/notifications";

import type { NotificationProps } from "@mantine/notifications";
import type { ReactNode } from "react";

type Props = {
  children: ReactNode;
  notificationProps: NotificationProps;
};

const NotificationButton = ({
  children,
  notificationProps,
}: Props): JSX.Element => {
  const notifications = useNotifications();

  return (
    <UnstyledButton
      onClick={() => notifications.showNotification(notificationProps)}
    >
      {children}
    </UnstyledButton>
  );
};

export default NotificationButton;
