import { PasswordInput as MantinePasswordInput } from "@mantine/core";
import { LockClosedIcon } from "@modulz/radix-icons";

import type { PasswordInputProps } from "@mantine/core/lib/components/PasswordInput/PasswordInput";
import type { UseForm } from "@mantine/hooks/lib/use-form/use-form";

export type Props<T extends { password: string }> = {
  form: UseForm<T>;
  props?: PasswordInputProps;
  disabled?: boolean;
};

const PasswordInput = <T extends { password: string }>({
  form,
  props,
  disabled,
}: Props<T>) => {
  return (
    <MantinePasswordInput
      required
      label="Password"
      placeholder="Password"
      toggleTabIndex={0}
      icon={<LockClosedIcon />}
      disabled={disabled}
      {...form.getInputProps(`password`)}
      {...props}
    />
  );
};

export default PasswordInput;
