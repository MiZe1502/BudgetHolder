import { Category, SimpleDataItem } from '../Categorization/types'
import { Shop } from '../Shops/types'
import { v4 as uuidv4 } from 'uuid'
import {SimpleReqDataItem} from "../../common/utils/api";

export class Purchase implements Purchase {
  constructor () {
    this.shop = {} as Shop
    this.goods = [
            {
              tempId: uuidv4(),
              category: {} as SimpleDataItem
            } as GoodsDetails
    ]
  }
}

export class GoodsItem implements Partial<GoodsData> {
    name = '';
    category = {} as SimpleDataItem;
    comment = '';

    constructor (newGoodsItem: GoodsData) {
      this.name = newGoodsItem.name
      this.category = newGoodsItem.category
      this.comment = newGoodsItem.comment
    }

    static fromGoodsDetails (item: GoodsDetails): GoodsItem {
      const goodsData = {
        name: item.name,
        category: item.category,
        comment: item.comment,
        id: item.id
      }
      return new GoodsItem(goodsData)
    }
}

export interface Purchase {
    id: number;
    totalPrice: number;
    comment?: string;
    shop: Shop;
    date: Date | string;
    goods: GoodsDetails[];
}

export interface GoodsDetails extends GoodsData, GoodsItemDetails {
    tempId?: string;
}

export interface GoodsData {
    id: number;
    name: string;
    category: SimpleDataItem | SimpleReqDataItem;
    comment?: string;
    category_id?: number;
}

export interface GoodsItemDetails {
    amount?: number;
    price: number;
    purchaseId: number;
}

export interface GoodsSuggestion {
    id: number;
    value: string;
}

export interface ValidationResult {
    message: string;
    goodsItemCounter?: number;
}
