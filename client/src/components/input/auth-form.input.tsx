import ConfirmPasswordInput from "@components/input/confirm-password.input";
import EmailInput from "@components/input/email.input";
import PasswordStrength from "@components/input/password-strength/password-strength.input";
import PasswordInput from "@components/input/password.input";
import UsernameInput from "@components/input/username.input";

import type { FormType } from "@custom-types/account.types";
import type { UseForm } from "@mantine/hooks/lib/use-form/use-form";
import type { FC } from "react";

type Props = {
  form: UseForm<{
    username: string;
    email: string;
    password: string;
    confirmPassword: string;
  }>;
  formType: FormType;
};

const AuthFormInput: FC<Props> = ({ form, formType }) => {
  return (
    <>
      {formType === `register` && <UsernameInput form={form} />}
      <EmailInput form={form} />
      {formType === `register` ? (
        <PasswordStrength form={form} />
      ) : (
        <PasswordInput form={form} />
      )}
      {formType === `register` && <ConfirmPasswordInput form={form} />}
    </>
  );
};

export default AuthFormInput;
