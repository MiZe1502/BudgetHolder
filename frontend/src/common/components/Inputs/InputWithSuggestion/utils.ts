export interface SuggestionItem {
    id: number;
    value: string;
}

export const maxSuggestionsListSize = 5;

export enum KeyboardKeys {
    ArrowUp = 38,
    ArrowDown = 40,
    Enter= 13,
}