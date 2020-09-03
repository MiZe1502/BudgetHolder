import {
    notInvalidByDefault,
    ValidationRule
} from "../../../common/utils/validation";
import {AuthData, RegistrationData} from "../../../stores/auth";

export const validationRules: ValidationRule[] = [
    {
        fieldName: "login",
        checkLength: true,
        length: 0,
        message: "Login field should not be empty",
        validator: function (data: AuthData | RegistrationData) {
            if (this.checkLength) {
                return data[this.fieldName].length === this.length;
            }
            return notInvalidByDefault;
        }
    },
    {
        fieldName: "password",
        checkLength: true,
        length: 0,
        message: "Password field should not be empty",
        validator: function (data: AuthData | RegistrationData) {
            if (this.checkLength) {
                return data[this.fieldName].length === this.length;
            }
            return notInvalidByDefault;
        }
    },
];
