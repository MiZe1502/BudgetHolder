import {
    notInvalidByDefault,
    ValidationRule,
} from "../../../common/utils/validation";
import {GoodsDetails, Purchase} from "../types";

export const validationRulesPurchase: ValidationRule[] = [
    {
        fieldName: "totalPrice",
        message: "Total price field should not be empty or lesser then 0",
        validator: function (data: Purchase) {
            return !Boolean(data[this.fieldName]) ||  data[this.fieldName] <= 0;
        }
    },
    {
        fieldName: "date",
        message: "Date field should not be empty",
        validator: function (data: Purchase) {
            return !Boolean(data[this.fieldName]);
        }
    },
    {
        fieldName: "shop",
        message: "Shop field should not be empty",
        validator: function (data: Purchase) {
            return !Boolean(data[this.fieldName]) || data[this.fieldName].id <= 0;
        }
    },
]

export const validationRulesGoods: ValidationRule[] = [
    {
        fieldName: "name",
        checkLength: true,
        length: 0,
        message: "Name field should not be empty",
        validator: function (data: GoodsDetails) {
            if (this.checkLength) {
                return data[this.fieldName].length === this.length;
            }
            return notInvalidByDefault;
        }
    },
    {
        fieldName: "category",
        message: "Category field should not be empty",
        validator: function (data: GoodsDetails) {
            return !Boolean(data[this.fieldName]) || data[this.fieldName].id <= 0;
        }
    },
    {
        fieldName: "amount",
        message: "Amount field should not be empty or lesser then 0",
        validator: function (data: GoodsDetails) {
            return !Boolean(data[this.fieldName]) ||  data[this.fieldName] <= 0;
        }
    },
    {
        fieldName: "price",
        message: "Price field should not be empty or lesser then 0",
        validator: function (data: GoodsDetails) {
            return !Boolean(data[this.fieldName]) ||  data[this.fieldName] <= 0;
        }
    },
]



