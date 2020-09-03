import {Category} from "../types";
import {
    notInvalidByDefault,
    ValidationRule,
} from "../../../common/utils/validation";

export const validationRules: ValidationRule[] = [
    {
        fieldName: "name",
        checkLength: true,
        length: 0,
        message: "Name field should not be empty",
        validator: function (data: Category) {
            if (this.checkLength) {
                return data[this.fieldName].length === this.length;
            }
            return notInvalidByDefault;
        }
    },
]



