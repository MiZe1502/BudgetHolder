import {Category, SimpleCategory} from "../Categorization/types";
import {Shop} from "../Shops/types";

export interface Purchase {
    id: number;
    totalPrice: number;
    comment?: string;
    shop: Shop;
    date: Date;
    goods: GoodsDetails[];
}

export interface GoodsDetails extends GoodsData, GoodsItemDetails {

}

export interface GoodsData {
    id: number;
    name: string;
    category: SimpleCategory;
    comment?: string;
}

export interface GoodsItemDetails {
    amount?: number;
    price: number;
}