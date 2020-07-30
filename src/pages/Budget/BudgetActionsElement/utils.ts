import {GoodsDetails} from "../types";

export interface BudgetActionsData {
    id: number;
    goodsData: GoodsDetails[];
    date: string;
    totalPrice: number;
}