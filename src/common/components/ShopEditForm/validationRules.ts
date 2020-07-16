import {Shop} from "../../../pages/types";
import {
    notInvalidByDefault,
    ValidationRule,
} from "../../utils/validation";

export const validationRules: ValidationRule[] = [
    {
        fieldName: "name",
        checkLength: true,
        length: 0,
        message: "Name field should not be empty",
        validator: function (data: Shop) {
            if (this.checkLength) {
                return data[this.fieldName].length === this.length;
            }
            return notInvalidByDefault;
        }
    },
    {
        fieldName: "name",
        checkLength: true,
        length: 100,
        message: "Name field is too long",
        validator: function (data: Shop) {
            if (this.checkLength) {
                return data[this.fieldName].length > this.length;
            }
            return notInvalidByDefault;
        }
    },
    {
        fieldName: "comment",
        checkLength: true,
        length: 1000,
        message: "Comment field is too long",
        validator: function (data: Shop) {
            if (this.checkLength) {
                return data[this.fieldName].length > this.length;
            }
            return notInvalidByDefault;
        }
    }
]



