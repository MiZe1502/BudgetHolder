import { Category, SimpleDataItem } from './types'

export const mockCategories: SimpleDataItem[] = [
  {
    id: 1,
    value: 'Products'
  },
  {
    id: 4,
    value: 'Milk'
  },
  {
    id: 5,
    value: 'Fish'
  },
  {
    id: 7,
    value: 'Snack'
  },
  {
    id: 8,
    value: 'Groceries'
  },
  {
    id: 2,
    value: 'Sport'
  },
  {
    id: 3,
    value: 'Entertainment'
  },
  {
    id: 6,
    value: 'Movies'
  },
  {
    id: 9,
    value: 'Games'
  },
  {
    id: 17,
    value: 'PS4'
  },
  {
    id: 18,
    value: 'PC'
  },
  {
    id: 10,
    value: 'Electronics'
  },
  {
    id: 11,
    value: 'TV'
  },
  {
    id: 12,
    value: 'Furniture'
  },
  {
    id: 13,
    value: 'Clothes'
  },
  {
    id: 14,
    value: 'Children'
  },
  {
    id: 15,
    value: 'Women'
  },
  {
    id: 16,
    value: 'Men'
  }

]

export const mockData: Category[] = [
  {
    id: 1,
    name: 'Products',
    parent_id: null,
    categories: [
      {
        id: 4,
        name: 'Milk',
        parent_id: 1
      },
      {
        id: 5,
        name: 'Fish',
        parent_id: 1
      },
      {
        id: 7,
        name: 'Snack',
        parent_id: 1
      },
      {
        id: 8,
        name: 'Groceries',
        parent_id: 1
      }
    ]
  },
  {
    id: 2,
    name: 'Sport',
    parent_id: null
  },
  {
    id: 3,
    name: 'Entertainment',
    parent_id: null,
    categories: [
      {
        id: 6,
        name: 'Movies',
        parent_id: 3
      },
      {
        id: 9,
        name: 'Games',
        parent_id: 3,
        categories: [
          {
            id: 17,
            name: 'PS4',
            parent_id: 9
          },
          {
            id: 18,
            name: 'PC',
            parent_id: 9
          }
        ]
      }
    ]
  },
  {
    id: 10,
    name: 'Electronics',
    parent_id: null,
    categories: [
      {
        id: 11,
        name: 'TV',
        parent_id: 10
      }
    ]
  },
  {
    id: 12,
    name: 'Furniture',
    parent_id: null
  },
  {
    id: 13,
    name: 'Clothes',
    parent_id: null,
    categories: [
      {
        id: 14,
        name: 'Children',
        parent_id: 13
      },
      {
        id: 15,
        name: 'Women',
        parent_id: 13
      },
      {
        id: 16,
        name: 'Men',
        parent_id: 13
      }
    ]
  }
]
