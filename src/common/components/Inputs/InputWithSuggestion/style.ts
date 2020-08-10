import {css} from "emotion";

export {
    Font312DarkGray, Font312RedAttention, Font312Black
} from "../../../styles/fonts";
export {FlexVert, FlexHorCenter} from "../../../styles/positioning";
export {InvalidInput, Input} from "../../../styles/inputs";

export const InputWithSuggestion = css`
  box-sizing: border-box;
  height: 40px;
  width: 100%;
`;

export const Wrapper = css`
  width: 100%;
  position: relative;
`;

export const DropdownInputExpanded = css`
    border-bottom: 0px !important;
    border-bottom-left-radius: 0px;
    border-bottom-right-radius: 0px;
`;