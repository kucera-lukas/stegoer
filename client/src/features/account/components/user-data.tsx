import { List, Text } from "@mantine/core";

import type { User } from "@graphql/generated/codegen.generated";

export type Props = {
  user: User;
};

const UserData = ({ user }: Props): JSX.Element => {
  return (
    <Text>
      Welcome {user.username}!
      <List>
        <List.Item>
          Last login date: {user.lastLogin.toLocaleString()}
        </List.Item>
        <List.Item>
          Account updated: {user.updatedAt.toLocaleString()}
        </List.Item>
        <List.Item>
          Account Created: {user.createdAt.toLocaleString()}
        </List.Item>
      </List>
    </Text>
  );
};

export default UserData;
