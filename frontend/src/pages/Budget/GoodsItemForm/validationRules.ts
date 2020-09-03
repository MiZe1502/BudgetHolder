import {
    notInvalidByDefault,
    ValidationRule
} from "../../../common/utils/validation";
import {GoodsData} from "../types";

export const validationRules: ValidationRule[] = [
    {
        fieldName: "name",
        checkLength: true,
        length: 0,
        message: "Name field should not be empty",
        validator: function (data: GoodsData) {
            if (this.checkLength) {
                return data[this.fieldName].length === this.length;
            }
            return notInvalidByDefault;
        }
    },
]
