import {css} from "emotion";

import {
    Font312DarkGray, Font312RedAttention, Font312Black
} from "../../../styles/fonts";
import {FlexVert, FlexHorCenter} from "../../../styles/positioning";
import {InvalidInput, Input} from "../../../styles/inputs";

export const style = {
    Font312DarkGray,
    Font312RedAttention,
    Font312Black,

    FlexVert,
    FlexHorCenter,

    InvalidInput,
    Input,

    InputWithSuggestion: css`
      box-sizing: border-box;
      height: 40px;
      width: 100%;
    `,

    Wrapper: css`
      width: 100%;
      position: relative;
    `,

    DropdownInputExpanded: css`
        border-bottom: 0px !important;
        border-bottom-left-radius: 0px;
        border-bottom-right-radius: 0px;
    `
}