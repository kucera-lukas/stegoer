import AccountNavigation from "@features/account/components/account.navigation";
import UpdateModal from "@features/account/components/modals/update.modal";
import UserData from "@features/account/components/user-data";

import { Title } from "@mantine/core";
import { useState } from "react";

import type { User } from "@graphql/generated/codegen.generated";
import type { FC } from "react";

type Props = {
  user: User;
};

const AccountComponent: FC<Props> = ({ user }) => {
  const [modelOpened, setModalOpened] = useState(false);

  return (
    <>
      <Title>Account</Title>
      <UserData user={user} />
      <UpdateModal
        user={user}
        opened={modelOpened}
        setOpened={setModalOpened}
      />
      <AccountNavigation
        user={user}
        disabled={modelOpened}
        onUpdate={() => setModalOpened(true)}
      />
    </>
  );
};

export default AccountComponent;