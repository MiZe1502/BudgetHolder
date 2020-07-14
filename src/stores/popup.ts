import { writable } from 'svelte/store';
import { get } from 'svelte/store';
import { v4 as uuidv4 } from 'uuid';
import error from "svelte/types/compiler/utils/error";

//TODO: We can implement this logic throught linked list

export const defaultPopupZIndex = 10;

export interface FormFieldErrors {
    fieldName: string;
    errors: string[];
}

export interface PopupState {
    zIndex?: number;
    uuid?: string;
    innerValidationErrors?: FormFieldErrors[];
    innerValidationFields?: Set<string>;
}
type PopupRecord = Record<string, PopupState>;

export const openedPopups = writable([]);

export const findCurrentTopPopup = (): PopupState => {
    const popups = get(openedPopups) as PopupRecord[];

    let topPopup: PopupState = popups[0];

    for (let i = 0; i < popups.length; i++) {
        if (popups[i].zIndex > topPopup.zIndex) {
            topPopup = popups[i];
        }
    }

    return topPopup;
}

export const addPopupToState = (uuid: string): PopupState => {
    const topPopup = findCurrentTopPopup();

    let curPopup: PopupState = {
        zIndex: defaultPopupZIndex, 
        uuid: uuid
    };

    if (topPopup) {
        curPopup.zIndex = topPopup.zIndex + 1;
    }

    openedPopups.update((popups) => [...popups, curPopup])

    return curPopup;
}

export const movePopupToTheTop = (uuid: string): PopupState => {
    let newTopPopup: PopupState = {};

    openedPopups.update((popups) => {
        const topPopup = findCurrentTopPopup();

        newTopPopup = popups.find((popup) => popup.uuid === uuid);
        newTopPopup.zIndex = topPopup.zIndex + 1;
        
        return popups;
    })

    return newTopPopup;
}

export const removePopupFromStore = (uuid: string) => {
    openedPopups.update((popups) => popups.filter((popup) => popup.uuid !== uuid));
}

//TODO: refactor needed
export const updatePopupInnerValidation = (uuid: string, errorMsg: string, fieldName: string) => {
    console.log("start", errorMsg, fieldName)
    openedPopups.update((popups) => {
        const popup = popups.find((popup) => popup.uuid === uuid);
        if (!popup.innerValidationErrors) {
            popup.innerValidationErrors = [{
                fieldName: fieldName,
                errors: [errorMsg],
            }]
        } else {
            const fieldErrors = popup.innerValidationErrors.find((error) => error.fieldName === fieldName);
            if (fieldErrors) {
                console.log(fieldErrors.errors.findIndex((error) => error === errorMsg))
                const isErrorMsgExists = fieldErrors.errors.findIndex((error) => error === errorMsg) !== -1;
                console.log("exists", isErrorMsgExists, fieldErrors)

                if (isErrorMsgExists) {
                    console.log("exists")
                } else {
                    popup.innerValidationErrors = [
                        ...popup.innerValidationErrors.filter((error) => error.fieldName !== fieldName),
                        {
                            fieldName: fieldName,
                            errors: [...fieldErrors.errors, errorMsg]
                        }
                    ];
                }
            } else {
                popup.innerValidationErrors = [
                    ...popup.innerValidationErrors,
                    {
                        fieldName: fieldName,
                        errors: [errorMsg]
                    }
                ];
            }
        }

        console.log(popups)

        return popups;
    })
}