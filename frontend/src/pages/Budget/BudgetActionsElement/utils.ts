import {GoodsDetails} from "../types";
import {Shop} from "../../Shops/types";

export interface BudgetActionsData {
    id: number;
    goods_data: GoodsDetails[];
    date: string;
    total_price: number;
    shop: Shop;
    shop_id: number;
}