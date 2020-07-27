import {Category} from "../Categorization/types";
import {Shop} from "../Shops/types";

export interface Purchase {
    totalPrice: number;
    comment?: string;
    shop: Shop;
    date: Date;
    goods: GoodsItem[];
}

export interface GoodsItem {
    id: number;
    name: string;
    category: Category;
    price: number;
    amount?: number;
    comment?: string;
}