import {css} from 'emotion'

import {
    Font312DarkGray, Font312RedAttention, Font312Black
} from "../../../styles/fonts";
import {FlexVert, FlexHorCenter} from "../../../styles/positioning";
import {InvalidInput, Input} from "../../../styles/inputs";

const maxHeight = 400;

export const style = {
    Font312DarkGray,
    Font312RedAttention,
    Font312Black,

    FlexVert,
    FlexHorCenter,

    InvalidInput,
    Input,

    maxHeight,

    Wrapper: css`
      width: 100%;
      position: relative;
    `,

    DropdownInput: css`
      cursor: pointer;
      width: 100%;
    `,

    DropdownInputExpanded: css`
        border-bottom: 0px !important;
        border-bottom-left-radius: 0px;
        border-bottom-right-radius: 0px;
    `,

    ButtonWrapper: css`
      position: absolute;
      right: 8px;
    `,

    IconButton: css`
      padding-top: 0px;
    `,

    ArrowIcon: css`
      & > svg {
        transform: rotate(90deg);
      }
    `,

    ArrowIconExpanded: css`
      & > svg {
        transform: rotate(270deg);
      }
    `
}