import {GoodsDetails} from "../types";
import {Shop} from "../../Shops/types";

export interface BudgetActionsData {
    id: number;
    goodsData: GoodsDetails[];
    date: string;
    totalPrice: number;
    shop: Shop;
}