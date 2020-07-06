import {css} from 'emotion';

import {defaultButtonIconSize} from "../../../styles/buttons";

export {
    ButtonReset,
    ButtonDefault,
    ButtonPrimary,
    ButtonSecondary,
    defaultButtonIconSize
} from "../../../styles/buttons";

export const Button = css`
    background-color: transparent;

    /* & > svg {
        transition: height 100ms, width 100ms;
    }

    &:hover > svg {
        width: ${defaultButtonIconSize + 2}px;
        height: ${defaultButtonIconSize + 2}px;
    } */

    &:hover,
	&:active,
    &:disabled {
		background-color: transparent;
	}
`;