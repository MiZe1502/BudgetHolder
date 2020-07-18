import {css} from 'emotion'

import colors from "../../../common/styles/colors";

export {Font312Black} from "../../../common/styles/fonts";
export {
    SideMinorPadding, FlexHorCenter
} from "../../../common/styles/positioning";

export const ArrowButton = css`
    margin-right: 8px;
`;

export const ArrowButtonExpanded = css`
    margin-right: 8px;

    & > svg {
        transform: rotate(90deg);
    }
`;

export const MarginedName = css`
    margin-left: 24px;
`;

export const CategoryLine = css`
    height: 40px;
    padding-right: 8px;
    border-bottom: 1px solid ${colors.darkGray};

    &:last-of-type {
        border-bottom: 0px;
    }
`;

export const CategoryLineExpanded = css`
  background-color: ${colors.gray};
`;