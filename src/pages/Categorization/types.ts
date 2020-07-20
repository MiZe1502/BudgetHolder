export interface Category {
    id: number;
    name: string;
    categories?: Category[];
    parentId: number | null;
}