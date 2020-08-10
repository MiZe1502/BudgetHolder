import {css} from 'emotion'

export {
    Font312DarkGray, Font312RedAttention, Font312Black
} from "../../../styles/fonts";
export {FlexVert, FlexHorCenter} from "../../../styles/positioning";
export {InvalidInput, Input} from "../../../styles/inputs";

export const maxHeight = 400;

export const Wrapper = css`
  width: 100%;
  position: relative;
`;

export const DropdownInput = css`
  cursor: pointer;
  width: 100%;
`;

export const DropdownInputExpanded = css`
    border-bottom: 0px !important;
    border-bottom-left-radius: 0px;
    border-bottom-right-radius: 0px;
`;

export const ButtonWrapper = css`
  position: absolute;
  right: 8px;
`;

export const IconButton = css`
  padding-top: 0px;
`;

export const ArrowIcon = css`
  & > svg {
    transform: rotate(90deg);
  }
`;

export const ArrowIconExpanded = css`
  & > svg {
    transform: rotate(270deg);
  }
`;